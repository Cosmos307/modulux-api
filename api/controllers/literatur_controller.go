package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"modulux/database"

	"github.com/gin-gonic/gin"
)

var crossRefURL string
var currentSuggestions []map[string]interface{}

// InitializeCrossRef initializes the CrossRef URL
func InitializeCrossRef(url string) {
	crossRefURL = url
}

// TestCrossRefConnection tests the connection to the CrossRef API
func TestCrossRefConnection(c *gin.Context) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/works?query.bibliographic=test&rows=2", crossRefURL), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request to CrossRef"})
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to CrossRef"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "CrossRef server returned an error",
			"status": resp.StatusCode,
			"body":   string(body),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully connected to CrossRef"})
}

// GetLiteratureSuggestions handles requests for literature suggestions
func GetLiteratureSuggestions(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter is required"})
		return
	}

	encodedQuery := url.QueryEscape(query)

	// request to CrossRef to get 5 suggestions
	requestURL := fmt.Sprintf("%s/works?query.bibliographic=%s&rows=5", crossRefURL, encodedQuery)
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request to CrossRef"})
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get suggestions from CrossRef"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response from CrossRef"})
		return
	}

	// Decode response from CrossRef
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode response from CrossRef"})
		return
	}

	message, ok := result["message"].(map[string]interface{})
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid response format"})
		return
	}

	items, ok := message["items"].([]interface{})
	if !ok || len(items) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No suggestions found"})
		return
	}

	var suggestions []map[string]interface{}
	for _, item := range items {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid item format"})
			return
		}

		titleArray, ok := itemMap["title"].([]interface{})
		if !ok || len(titleArray) == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid title format"})
			return
		}
		title := strings.Join(convertInterfaceArrayToStringArray(titleArray), " ")

		authors, ok := itemMap["author"].([]interface{})
		if !ok {
			authors = []interface{}{}
		}
		var formattedAuthors []string
		for _, author := range authors {
			authorMap, ok := author.(map[string]interface{})
			if ok {
				formattedAuthors = append(formattedAuthors, fmt.Sprintf("%s %s", authorMap["given"], authorMap["family"]))
			}
		}
		authorStr := strings.Join(formattedAuthors, ", ")

		publishedArray, ok := itemMap["published"].(map[string]interface{})["date-parts"].([]interface{})
		if !ok || len(publishedArray) == 0 || len(publishedArray[0].([]interface{})) == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid published format"})
			return
		}
		year := int(publishedArray[0].([]interface{})[0].(float64))

		publisher, ok := itemMap["publisher"].(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid publisher format"})
			return
		}
		doi, ok := itemMap["DOI"].(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid DOI format"})
			return
		}
		url, ok := itemMap["URL"].(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid URL format"})
			return
		}

		// Add suggestion to the list
		suggestion := map[string]interface{}{
			"title":     title,
			"author":    authorStr,
			"year":      year,
			"publisher": publisher,
			"doi":       doi,
			"url":       url,
		}
		suggestions = append(suggestions, suggestion)
	}

	currentSuggestions = suggestions

	c.JSON(http.StatusOK, suggestions)
}

// Convert an array of interface{} to an array of strings
func convertInterfaceArrayToStringArray(arr []interface{}) []string {
	var strArr []string
	for _, v := range arr {
		strArr = append(strArr, v.(string))
	}
	return strArr
}

// ConfirmLiteratureSelection handles the confirmation of the user's literature selection
func ConfirmLiteratureSelection(c *gin.Context) {
	var selection struct {
		ModuleKuerzel string `json:"module_kuerzel"`
		ModuleVersion int    `json:"module_version"`
		DOI           string `json:"doi"`
	}
	if err := c.ShouldBindJSON(&selection); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Check if the selected suggestion is in the current suggestions
	var selectedSuggestion map[string]interface{}
	for _, suggestion := range currentSuggestions {
		if suggestion["doi"] == selection.DOI {
			selectedSuggestion = suggestion
			break
		}
	}

	if selectedSuggestion == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Selected suggestion is not in the current suggestions"})
		return
	}

	tx, err := database.DB.Begin(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to begin transaction"})
		return
	}
	defer tx.Rollback(context.Background())

	// Check if the DOI already exists in the database
	var existingLiteraturID int
	var literaturID int
	err = tx.QueryRow(context.Background(), `
		SELECT literatur_id
		FROM literatur
		WHERE doi = $1
	`, selection.DOI).Scan(&existingLiteraturID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// DOI does not exist, insert the new literature
			query := `
				INSERT INTO literatur (titel, autor, jahr, verlag, isbn, link, doi)
				VALUES ($1, $2, $3, $4, $5, $6, $7)
				RETURNING literatur_id
			`
			err = tx.QueryRow(context.Background(), query,
				selectedSuggestion["title"],
				selectedSuggestion["author"],
				selectedSuggestion["year"],
				selectedSuggestion["publisher"],
				nil, // ISBN is not provided by CrossRef
				selectedSuggestion["url"],
				selectedSuggestion["doi"],
			).Scan(&literaturID)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert literature into database"})
				return
			}
		} else {
			fmt.Println("Error checking existing DOI:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing DOI", "details": err.Error()})
			return
		}
	} else {
		// DOI exists, use the existing literatur_id
		literaturID = existingLiteraturID
	}

	fmt.Println("Literatur ID durch:", literaturID)

	// Ermitteln der höchsten snapshot_id für das spezifische Modul und die Version
	var maxSnapshotID sql.NullInt32
	err = tx.QueryRow(context.Background(), `
        SELECT COALESCE(MAX(snapshot_id), 0)
        FROM modul_literatur_historie
        WHERE modul_kuerzel = $1 AND modul_version = $2
    `, selection.ModuleKuerzel, selection.ModuleVersion).Scan(&maxSnapshotID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get snapshot IDs"})
		return
	}
	fmt.Println("Max Snapshot ID:", maxSnapshotID)
	newSnapshotID := maxSnapshotID.Int32 + 1

	// Check if there are existing literature references for the module and version
	var count int
	err = tx.QueryRow(context.Background(), `
        SELECT COUNT(*)
        FROM modul_literatur
        WHERE modul_kuerzel = $1 AND modul_version = $2
    `, selection.ModuleKuerzel, selection.ModuleVersion).Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing literature"})
		return
	}

	var query string
	fmt.Println("count:", count)
	if count > 0 {
		// insert the existing literature references into modul_literatur_historie
		_, err = tx.Exec(context.Background(), `
            INSERT INTO modul_literatur_historie (modul_kuerzel, modul_version, literatur_id, snapshot_id, vorheriger_snapshot_id, aenderungsdatum)
            SELECT modul_kuerzel, modul_version, literatur_id, $1, 
            CASE WHEN $2 = 0 THEN NULL ELSE $2 END, NOW()
            FROM modul_literatur
            WHERE modul_kuerzel = $3 AND modul_version = $4
        `, newSnapshotID, maxSnapshotID.Int32, selection.ModuleKuerzel, selection.ModuleVersion)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save literature history"})
			return
		}

		updateQuery := `
			UPDATE modul_literatur
			SET vorheriger_snapshot_id = $1
			WHERE modul_kuerzel = $2 AND modul_version = $3
		`
		_, err = tx.Exec(context.Background(), updateQuery, newSnapshotID, selection.ModuleKuerzel, selection.ModuleVersion)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update previous snapshot"})
			return
		}

		query = `
			INSERT INTO modul_literatur (modul_kuerzel, modul_version, literatur_id, vorheriger_snapshot_id)
			VALUES ($1, $2, $3, $4)
   		`
		_, err = tx.Exec(context.Background(), query, selection.ModuleKuerzel, selection.ModuleVersion, literaturID, newSnapshotID)
		if err != nil {
			fmt.Println("Error inserting relation:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert relation into database"})
			return
		}
	} else {
		query = `
            INSERT INTO modul_literatur (modul_kuerzel, modul_version, literatur_id, vorheriger_snapshot_id)
            VALUES ($1, $2, $3, NULL)
        `
		_, err = tx.Exec(context.Background(), query, selection.ModuleKuerzel, selection.ModuleVersion, literaturID)
		if err != nil {
			fmt.Println("Error inserting relation:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert relation into database"})
			return
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Selection confirmed", "selected_suggestion": selectedSuggestion})
}
