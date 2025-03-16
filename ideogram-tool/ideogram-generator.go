package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Config struct {
	ApiKey            string `json:"api_key"`
	Prompt            string `json:"prompt"`
	NumImages         int    `json:"num_images"`
	OutputPath        string `json:"output_path"`
	BaseImageName     string `json:"base_image_name"`
	AspectRatio       string `json:"aspect_ratio"`
	StyleType         string `json:"style_type"`
	MagicPromptOption string `json:"magic_prompt_option"`
	NumberPadding     int    `json:"number_padding"`
	NumberSeparator   string `json:"number_separator"`
	StartNumber       int    `json:"start_number"`
}

type IdeogramResponse struct {
	Created string `json:"created"`
	Data    []struct {
		URL string `json:"url"`
	} `json:"data"`
}

func main() {
	configPath := flag.String("config", "config.json", "Percorso del file di configurazione")
	flag.Parse()

	config, err := loadConfig(*configPath)
	if err != nil {
		fmt.Printf("Errore nella lettura del config %s: %v\n", *configPath, err)
		return
	}

	if config.NumberPadding <= 0 {
		config.NumberPadding = 3
	}
	if config.NumberSeparator == "" {
		config.NumberSeparator = "_"
	}
	if config.StartNumber < 0 {
		config.StartNumber = 1
	}

	err = os.MkdirAll(config.OutputPath, 0755)
	if err != nil {
		fmt.Printf("Errore nella creazione della directory: %v\n", err)
		return
	}

	for i := 0; i < config.NumImages; i++ {
		currentNumber := config.StartNumber + i
		fmt.Printf("Generazione immagine %d di %d...\n", i+1, config.NumImages)

		resp, err := generateImage(config)
		if err != nil {
			fmt.Printf("Errore nella generazione dell'immagine %d: %v\n", currentNumber, err)
			continue
		}

		err = downloadImage(resp.Data[0].URL, config, currentNumber)
		if err != nil {
			fmt.Printf("Errore nel download dell'immagine %d: %v\n", currentNumber, err)
			continue
		}

		if i < config.NumImages-1 {
			time.Sleep(2 * time.Second)
		}
	}

	fmt.Println("Generazione completata!")
}

func loadConfig(filename string) (*Config, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func generateImage(config *Config) (*IdeogramResponse, error) {
	url := "https://api.ideogram.ai/generate"

	requestBody := fmt.Sprintf(`{
		"image_request": {
			"prompt": "%s",
			"aspect_ratio": "%s",
			"model": "V_2",
			"magic_prompt_option": "%s",
			"style_type": "%s"
		}
	}`, config.Prompt, config.AspectRatio, config.MagicPromptOption, config.StyleType)

	payload := strings.NewReader(requestBody)
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Api-Key", config.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response IdeogramResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func downloadImage(imageUrl string, config *Config, number int) error {
	response, err := http.Get(imageUrl)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	paddedNumber := fmt.Sprintf("%0*d", config.NumberPadding, number)
	filename := filepath.Join(
		config.OutputPath,
		fmt.Sprintf("%s%s%s.png", config.BaseImageName, config.NumberSeparator, paddedNumber),
	)

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Immagine salvata: %s\n", filename)
	return nil
}
