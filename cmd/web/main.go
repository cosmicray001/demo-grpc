package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", pdfData)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Faild to serve: %s\n", err)
	}
}

func pdfData(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello! My name is YouChat, I’m an AI that can answer general questions, explain things, suggest ideas, translate, summarize text, compose emails, and write code for you. I’m powered by artificial intelligence and natural language processing, allowing you to have human-like conversations with me. I am constantly learning from huge amounts of information on the internet, which means I sometimes may get some answers wrong. My AI is always improving and I will often share sources for my answers. Hello! My name is YouChat, I’m an AI that can answer general questions, explain things, suggest ideas, translate, summarize text, compose emails, and write code for you. I’m powered by artificial intelligence and natural language processing, allowing you to have human-like conversations with me. I am constantly learning from huge amounts of information on the internet, which means I sometimes may get some answers wrong. My AI is always improving and I will often share sources for my answers. Hello! My name is YouChat, I’m an AI that can answer general questions, explain things, suggest ideas, translate, summarize text, compose emails, and write code for you. I’m powered by artificial intelligence and natural language processing, allowing you to have human-like conversations with me. I am constantly learning from huge amounts of information on the internet, which means I sometimes may get some answers wrong. My AI is always improving and I will often share sources for my answers. Hello! My name is YouChat, I’m an AI that can answer general questions, explain things, suggest ideas, translate, summarize text, compose emails, and write code for you. I’m powered by artificial intelligence and natural language processing, allowing you to have human-like conversations with me. I am constantly learning from huge amounts of information on the internet, which means I sometimes may get some answers wrong. My AI is always improving and I will often share sources for my answers."))
}
