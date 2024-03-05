package servicos

import (
	"bytes"
	"encoding/json"
	"fmt"
	"http_gin/src/model_views"
	"io"
	"net/http"
	"os"
)

type IAServico struct{}

func (pc *IAServico) BuscarPalavras() []model_views.Palavra {

	// respostaTeste := `
	// [
	// 	{"palavra": "played", "traducao": "jogou", "opcoes": ["dançou", "cantou", "correu", "jogou"] },
	// 	{"palavra": "visited", "traducao": "visitou", "opcoes": ["estudou", "assistiu", "cozinhou", "visitou"] },
	// 	{"palavra": "studied", "traducao": "estudou", "opcoes": ["trabalhou", "brincou", "correu", "estudou"] },
	// 	{"palavra": "listened", "traducao": "ouviu", "opcoes": ["falou", "gritou", "correu", "ouviu"] }
	// ]
	// `

	// var palavrasTest []model_views.Palavra
	// err := json.Unmarshal([]byte(respostaTeste), &palavrasTest)
	// if err != nil {
	// 	fmt.Println("Erro ao deserializar resposta:", err)
	// 	return nil
	// }
	// return palavrasTest

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("OPENAI_API_KEY não definida.")
		return nil
	}

	messages := []map[string]string{
		{
			"role":    "system",
			"content": "Sua missão é retornar para mim uma lista de palavras em ingles com quatro alternativas e uma correta, o formato retornado será assim: [{\"palavra\": \"Hello\", \"traducao\": \"Olá\", opcoes: [\"Boa\", \"Ok\", \"Olá\", \"Bacana\"] }] somente o json e nenhum outro texto",
		},
		{
			"role":    "user",
			"content": "Traga para mim palavras em ingles no formato passado, as opções retornam em pt-br",
		},
	}

	url := "https://api.openai.com/v1/chat/completions"

	requestBody, err := json.Marshal(map[string]interface{}{
		"model":    "gpt-3.5-turbo",
		"messages": messages,
	})

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	responseBody, erro := io.ReadAll(resp.Body)

	if erro != nil {
		fmt.Println(err.Error())
		return nil
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(responseBody), &result)

	var palavras []model_views.Palavra

	fmt.Println("=============")
	fmt.Println(result)
	fmt.Println("=============")

	if choices, ok := result["choices"].([]interface{}); ok && len(choices) > 0 {
		if firstChoice, ok := choices[0].(map[string]interface{}); ok {
			if message, ok := firstChoice["message"].(map[string]interface{}); ok {
				resposta := message["content"].(string)

				fmt.Println("=============")
				fmt.Println(resposta)
				fmt.Println("=============")

				err := json.Unmarshal([]byte(resposta), &palavras)
				if err != nil {
					fmt.Println("Erro ao deserializar resposta:", err)
					return nil
				}
				return palavras
			}
		}
	}

	fmt.Println("Não foi possível obter uma resposta válida.")
	return nil
}
