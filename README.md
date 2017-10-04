# Simple-Score-API

Simple example to learn golang

Starts a http server on port 8000

Get current scores with a GET request "/score"
Add new score (overwrites oldscore with same name) with POST request "/score" and JSON Body {"name": "<your name>", "score": 100}

Based on tutorial from https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo
