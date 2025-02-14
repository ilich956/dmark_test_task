import { writable, get } from 'svelte/store';
import {GetTasks, InsertTask, UpdateTask, DeleteTask } from "../wailsjs/go/backend/App.js";

//store for list of tasks
export let tasks = writable([]);
//store for new tasks
export  let newTask = writable({ title: '', description: '', deadline: '', priority: '', status: '' });
//control alert visibility
export  let isAlert = writable(false);

/**
 * Fetch tasks from the backend and update the `tasks` store
 */
export async function fetchTasks() {
    try {
      const result = await GetTasks(); // Fetch tasks from backend
      if (result == null) {
        tasks.set([]);   //if no tasks, set an empty array
      } else {
        tasks.set(result); //update store with fetched tasks
      }

      console.log("fetched: ", result);
    } catch (error) {
      alert("Error fetching tasks:")
      console.error("Error fetching tasks:", error);
    }
  }

 /**
 * Add a new task to the backend and refresh the task list
 */ 
export async function addTask() {
    try {
      const task = get(newTask);
      if (task.title == "") {  // validate task title
        isAlert.set(true);
        setTimeout(() => {
          isAlert.set(false); // hide alert after 3 seconds
        }, 3000);

        return;
      }

      // set default values if not provided
      if (task.priority == "") {
        task.priority = "Low";
      }
      if (task.status == "") {
        task.status = "Pending";
      }
      if (task.deadline == "") {
        task.deadline = "none";
      }
      // insert the new task into the backend
      await InsertTask(task.title, task.description, task.deadline, task.priority, task.status);

       // reset the form fields
      newTask.set({ title: "", description: "", deadline: "", priority: "", status: "" });
      // refresh task list
      fetchTasks();
    } catch (error) {
      alert("Error adding task:");
      console.error("Error adding task:", error);
    }
  }

  /**
 * Delete a task from the backend and refresh the task list.
 */
export async function deleteTask(id) {
    try {
      await DeleteTask(id);
      fetchTasks();
    } catch (error) {
      alert("Error deleting task:");
      console.error("Error deleting task:", error);
    }
  }

  /**
 * Update the priority and status of a task in the backend and refresh the task list.
 */
export async function updateTask(id, priority, status) {
    try {
      await UpdateTask(id, priority, status);
      fetchTasks();
    } catch (error) {
      alert("Error updating task:");
      console.error("Error updating task:", error);
    }
  }