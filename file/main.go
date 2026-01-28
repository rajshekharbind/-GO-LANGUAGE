package main
import (
	"fmt"
	"os"
	
)
func main() {
	fmt.Println("i want to learn consepts of file in golang:")

	filePath := "example.txt"

	// Creating a file
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	fmt.Println("File created successfully:", filePath)

	// Writing to the file
	content := "Hello, this is a sample text written to the file."
	_, err = file.WriteString(content)

	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Content written to file successfully.")

	// Reading from the file
	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer readFile.Close()


	fileInfo, err := readFile.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	_, err = readFile.Read(buffer)

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Content read from file successfully:")
	fmt.Println(string(buffer))

	// Deleting the file
	err = os.Remove(filePath)
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}
	fmt.Println("File deleted successfully:", filePath)
     
	// Getting current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}
	fmt.Println("Current working directory:", cwd)

	// Listing files in the current directory
	files, err := os.ReadDir(cwd)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}
	fmt.Println("Files in current directory:")
	for _, file := range files {
		fmt.Println(" -", file.Name())
	}

	// Renaming a file (creating a new file first)
	newFilePath := "renamed_example.txt"
	_, err = os.Create(newFilePath)

	if err != nil {
		fmt.Println("Error creating file for renaming:", err)
		return
	}
	err = os.Rename(newFilePath, "final_example.txt")
	if err != nil {
		fmt.Println("Error renaming file:", err)
		return
	}
	fmt.Println("File renamed successfully to final_example.txt")

	// Getting file permissions
	finalFileInfo, err := os.Stat("final_example.txt")
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}
	fmt.Println("File permissions of final_example.txt:", finalFileInfo.Mode())

	// Clean up
	err = os.Remove("final_example.txt")
	if err != nil {
		fmt.Println("Error deleting final_example.txt:", err)
		return
	}
	fmt.Println("Cleaned up final_example.txt")

	// Environment variables
	fmt.Println("Environment Variables:")
	for _, env := range os.Environ() {
		fmt.Println(" -", env)
	}

	// Creating a directory
	dirPath := "sampleDir"
	err = os.Mkdir(dirPath, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
	fmt.Println("Directory created successfully:", dirPath)

	// Removing the directory
	err = os.Remove(dirPath)
	if err != nil {
		fmt.Println("Error removing directory:", err)
		return
	}
	fmt.Println("Directory removed successfully:", dirPath)

	// Exiting the program with a specific exit code
	fmt.Println("Exiting the program with exit code 0.")
	os.Exit(0)

	// Note: Code after os.Exit will not be executed
	fmt.Println("This line will not be printed.")

	// End of main function
	
}