# dotenv file Parser

This repository implements a .env file parser, handling invalid .env formats

## In this README 👇

- [Features](#features)
- [Usage](#usage)

## Features
 - `Parse(path string)` takes the .env file path, loads and parses them into a map of key-value pairs

## Usage

1. 
    ```go
        import github.com/codescalersinternships/dotenv-rawan
    ```


2. Example usage:
    ```go
        envs, err := Parse(filePath)
			if err != nil {
				return err
		}
    ```


