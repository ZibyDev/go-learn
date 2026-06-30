package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Reading strings
var reader = bufio.NewReader(os.Stdin)

func readLine(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// Commands type
type Command int

const (
	show = iota + 1
	add
	del
	complete
	filter
	stats
	search
	exit
)

var stringToCommand = map[string]Command{
	"show":     show,
	"add":      add,
	"delete":   del,
	"complete": complete,
	"filter":   filter,
	"stats":    stats,
	"search":   search,
	"exit":     exit,
}

// Filter commands type
type FilterCommand int

const (
	fComplete = iota + 1
	fUncomplete
	fPriority
)

var stringToFilterCommand = map[string]FilterCommand{
	"complete":   fComplete,
	"uncomplete": fUncomplete,
	"priority":   fPriority,
}

// Priority type
type Priority int

const (
	Low Priority = iota + 1
	Medium
	High
)

var stringToPriority = map[string]Priority{
	"low":    Low,
	"medium": Medium,
	"high":   High,
}

func (p Priority) String() string {
	switch p {
	case Low:
		return "Low"
	case Medium:
		return "Medium"
	case High:
		return "High"
	}
	return "Unknown"
}

// Task struct
type Task struct {
	ID          int
	TaskName    string
	Description string
	Priority    Priority
	IsCompleted bool
}

func (t Task) String() string {
	return fmt.Sprintf("ID: %v\tName: %v\tPriority: %v\tIsCompleted: %v\tDescription: %v", t.ID, t.TaskName, t.Priority, t.IsCompleted, t.Description)
}

// Functions to work with menu
func printMenu() {
	fmt.Println("\nAvailable commands:")
	fmt.Println("show\t\tShow all tasks")
	fmt.Println("add\t\tAdd task")
	fmt.Println("delete\t\tDelete task")
	fmt.Println("complete\tComplete task")
	fmt.Println("filter\t\tFilter tasks")
	fmt.Println("stats\t\tShow statistics")
	fmt.Println("search\t\tSearch task")
	fmt.Println("exit\t\tSave and exit")
}

func inputCommand(s string) Command {
	return stringToCommand[readLine(s)]
}

func inputFilterCommand(s string) FilterCommand {
	return stringToFilterCommand[readLine(s)]
}

func inputPriority(s string) Priority {
	return stringToPriority[readLine(s)]
}

// CRUD functions
func showTasks(tasksList map[int]Task) {
	fmt.Println("\nYour tasks:")
	for _, task := range tasksList {
		fmt.Println(task)
	}
}

func addTask(tasksList map[int]Task, totalTasks *int) {
	newTask := Task{ID: *totalTasks, IsCompleted: false}

	taskName := readLine("\nEnter task name: ")
	for taskName == "" {
		taskName = readLine("Task name can't be empty. Enter task name: ")
	}
	newTask.TaskName = taskName

	prio := inputPriority("\nEnter task priority (low, medium, high): ")
	for prio < 1 || prio > 3 {
		prio = inputPriority("Unknown priority. Enter task priority (low, medium, high): ")
	}
	newTask.Priority = Priority(prio)

	newTask.Description = readLine("Enter task description (optional): ")

	tasksList[*totalTasks] = newTask
	*totalTasks++
}

func deleteTask(tasksList map[int]Task) {
	id, err := strconv.Atoi(readLine("\nEnter task ID you want to delete: "))
	if err != nil {
		fmt.Println("ID must be a number")
		return
	}
	if _, ok := tasksList[id]; ok {
		delete(tasksList, id)
		fmt.Println("Successful task deletion")
	} else {
		fmt.Printf("Task with ID %d does not exist\n", id)
	}
}

func completeTask(tasksList map[int]Task) {
	id, err := strconv.Atoi(readLine("\nEnter task ID you want to complete: "))
	if err != nil {
		fmt.Println("ID must be a number")
		return
	}
	task, ok := tasksList[id]
	if ok {
		task.IsCompleted = true
		tasksList[id] = task
		fmt.Println("Successful task completion")
	} else {
		fmt.Printf("Task with ID: %v does not exist\n", id)
	}
}

// Filter functions
func printFilterMenu() {
	fmt.Println("\nAvailable filter commands:")
	fmt.Println("complete - Show completed")
	fmt.Println("uncomplete - Show uncompleted")
	fmt.Println("priority - Show by Priority")
}

func filterComplete(tasksList map[int]Task) {
	fmt.Println("\nCompleted tasks:")
	for _, task := range tasksList {
		if task.IsCompleted {
			fmt.Println(task)
		}
	}
}
func filterUncomplete(tasksList map[int]Task) {
	fmt.Println("\nUncompleted tasks:")
	for _, task := range tasksList {
		if !task.IsCompleted {
			fmt.Println(task)
		}
	}
}
func filterByPriority(tasksList map[int]Task) {
	prio := inputPriority("\nEnter task priority to filter (low, medium, high): ")
	if prio > 0 && prio < 4 {
		fmt.Printf("\n%v priority tasks:\n", Priority(prio))
		for _, task := range tasksList {
			if task.Priority == Priority(prio) {
				fmt.Println(task)
			}
		}
	} else {
		fmt.Println("Unknown priority")
	}
}

func filterTasks(tasksList map[int]Task) {
	printFilterMenu()
	choice := inputFilterCommand("Your filter command: ")
	switch choice {
	case fComplete:
		filterComplete(tasksList)
	case fUncomplete:
		filterUncomplete(tasksList)
	case fPriority:
		filterByPriority(tasksList)
	default:
		fmt.Println("Unknown command")
	}
}

// Stats function
func showStatistics(tasksList map[int]Task) {
	var all, completed, low, medium, high int
	for _, task := range tasksList {
		all++
		if task.IsCompleted {
			completed++
		}
		switch task.Priority {
		case Low:
			low++
		case Medium:
			medium++
		case High:
			high++
		}
	}
	fmt.Printf("\nYou have:\n%v task total\n%v task complete\n%v low priority tasks\n%v medium priority tasks\n%v high priority tasks\n", all, completed, low, medium, high)
}

// Search function
func searchTasks(tasksList map[int]Task) {
	searchText := strings.ToLower(readLine("\nEnter string you want to search: "))
	fmt.Println("Found tasks:")
	for _, task := range tasksList {
		if strings.Contains(strings.ToLower(task.TaskName), searchText) {
			fmt.Println(task)
		}
	}
}

// Save and load functions
func saveToFile(filename string, tasks map[int]Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func loadFromFile(filename string, totalTasks *int) (map[int]Task, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return map[int]Task{}, nil
		}
		return nil, err
	}
	var tasks map[int]Task
	err = json.Unmarshal(data, &tasks)
	if len(tasks) != 0 {
		for i := range tasks {
			if i > *totalTasks {
				*totalTasks = i
			}
		}
		*totalTasks++
	}
	return tasks, err
}

// Main cycle for Task Manager
func main() {
	var totalTasks = 0
	var tasksList, _ = loadFromFile("saveTasks.json", &totalTasks)
mainCycle:
	for {
		printMenu()
		choice := inputCommand("Your command: ")
		switch choice {
		case show:
			showTasks(tasksList)
		case add:
			addTask(tasksList, &totalTasks)
		case del:
			deleteTask(tasksList)
		case complete:
			completeTask(tasksList)
		case filter:
			filterTasks(tasksList)
		case stats:
			showStatistics(tasksList)
		case search:
			searchTasks(tasksList)
		case exit:
			break mainCycle
		default:
			fmt.Println("Unknown command")
		}
	}
	saveToFile("saveTasks.json", tasksList)
}
