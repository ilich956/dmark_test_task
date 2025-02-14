import { writable, get } from 'svelte/store';
import {GetTasks, InsertTask, UpdateTask, DeleteTask } from "../wailsjs/go/backend/App.js";

export let tasks = writable([]);
export  let newTask = writable({ title: '', description: '', deadline: '', priority: '', status: '' });
export  let isAlert = writable(false);

export async function fetchTasks() {
    try {
      const result = await GetTasks();
      if (result == null) {
        tasks.set([]);
      } else {
        tasks.set(result);
      }

      console.log("fetched: ", result);
    } catch (error) {
      alert("Error fetching tasks:")
      console.error("Error fetching tasks:", error);
    }
  }

export async function addTask() {
    try {
      const task = get(newTask);
      if (task.title == "") {
        isAlert.set(true);
        setTimeout(() => {
          isAlert.set(false);
        }, 3000);

        return;
      }
      if (task.priority == "") {
        task.priority = "Low";
      }
      if (task.status == "") {
        task.status = "Pending";
      }
      if (task.deadline == "") {
        task.deadline = "none";
      }

      await InsertTask(task.title, task.description, task.deadline, task.priority, task.status);
      newTask.set({ title: "", description: "", deadline: "", priority: "", status: "" });
      fetchTasks();
    } catch (error) {
      alert("Error adding task:");
      console.error("Error adding task:", error);
    }
  }

export async function deleteTask(id) {
    try {
      await DeleteTask(id);
      fetchTasks();
    } catch (error) {
      alert("Error deleting task:");
      console.error("Error deleting task:", error);
    }
  }

export async function updateTask(id, priority, status) {
    try {
      await UpdateTask(id, priority, status);
      fetchTasks();
    } catch (error) {
      alert("Error updating task:");
      console.error("Error updating task:", error);
    }
  }