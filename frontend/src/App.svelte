<script>
import { onMount } from 'svelte';
import {SvelteUIProvider , Button, Input, Textarea, Card, Title, Group, Badge, NativeSelect, Alert, Modal } from '@svelteuidev/core';
import { ChevronDown, InfoCircled } from 'radix-icons-svelte';
import { writable } from 'svelte/store';
import { GetTasks, InsertTask, UpdateTask, DeleteTask } from "../wailsjs/go/backend/App.js";

  let tasks = writable([]);
  let newTask = writable({ title: '', description: '', deadline: '', priority: '', status: '' });


  let showDeletePopUp = writable(false);
  let showUpdatePopUp = writable(false);

  let taskToDelete = writable(null);
  let taskToUpdate = writable(null);

  let isAlert = writable(false);

  let sortOrder = writable("ascending");
  const priorityOrder = {
        "Low": 1,
        "Medium": 2,
        "High": 3
    };

  onMount(async () => {
    fetchTasks();
  });

  async function fetchTasks() {
    try {
      const result = await GetTasks();
      if (result == null) {
        tasks.set([]);
      } else {
        tasks.set(result);
      }

      console.log(result);
    } catch (error) {
      console.error("Error fetching tasks:", error);
    }
  }

  function sortTasksByPriority() {
    tasks.update(currentTasks => {
        const order = $sortOrder === "ascending" ? 1 : -1;
        return [...currentTasks].sort((a, b) => (priorityOrder[a.Priority] - priorityOrder[b.Priority]) * order);
    });
}


  async function addTask() {
    const task = $newTask;
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

    await InsertTask( task.title, task.description, task.deadline, task.priority, task.status);
    newTask.set({ title: "", description: "", deadline: "", priority: "", status: "",});
    fetchTasks();
  }

  async function deleteTask(id) {
    await DeleteTask(id);
    fetchTasks();
  }

  async function updateTask(id, updatedTask) {
    await UpdateTask(id, updatedTask.priority, updatedTask.status);
    fetchTasks();
  }

  async function completeTask(task) {
    await UpdateTask(task.ID, task.Priority, "Completed");
    fetchTasks();
  }

  async function uncompleteTask(task) {
    await UpdateTask(task.ID, task.Priority, "In progress");
    fetchTasks();
  }

  function openDeletePopUp(task) {
    taskToDelete.set(task.ID);
    showDeletePopUp.set(true);
  }

  function openUpdatePopUp(task) {
    console.log("task: up: ", task);
    taskToUpdate.set({
      ID: task.ID,
      priority: task.Priority,
      status: task.Status,
    });
    showUpdatePopUp.set(true);
  }

  function closeDeletePopUp() {
    showDeletePopUp.set(false);
  }

  function closeUpdatePopUp() {
    showUpdatePopUp.set(false);
  }

  async function confirmDelete() {
    const id = $taskToDelete;
    deleteTask(id);
    closeDeletePopUp();
  }

  async function confirmUpdate() {
    const task = $taskToUpdate;
    updateTask(task.ID, task);
    closeUpdatePopUp();
  }

  $: completedTasks = $tasks.filter((task) => task.Status === "Completed");
  $: console.log($tasks);
</script>

<SvelteUIProvider>
  <main>
    {#if $isAlert}
      <Alert icon={InfoCircled} title="Oopsie!" variant="filled" radius="xs" withCloseButton on:close={() => isAlert.set(false)}>
        Enter a title for the task!
      </Alert>
    {/if}
    <Title order={1} align="center" style="color: aliceblue; padding-bottom: 20px;">Simple todo list 📝</Title
    >

    <Card withBorder shadow="md" padding="xl" mt="md">
      <Group position="apart" mb="md">
        <Input placeholder="Task title" bind:value={$newTask.title} />
        <Input
          placeholder="Deadline"
          type="date"
          bind:value={$newTask.deadline}
        />
        <NativeSelect
          placeholder="Priority"
          bind:value={$newTask.priority}
          data={["Low", "Medium", "High"]}
        >
          <svelte:component this={ChevronDown} slot="rightSection" />
        </NativeSelect>
        <NativeSelect
          placeholder="Status"
          bind:value={$newTask.status}
          data={["Pending", "In progress", "Completed"]}
        >
          <svelte:component this={ChevronDown} slot="rightSection" />
        </NativeSelect>
      </Group>
      <Textarea
        placeholder="Task description"
        bind:value={$newTask.description}
        mb="md"
      />

      <Group position="apart">
      <Button on:click={addTask}>Add Task</Button>
      <NativeSelect bind:value={$sortOrder} data={["ascending", "descending"]} on:change={sortTasksByPriority} label="Sort by priority" >
        <svelte:component this={ChevronDown} slot="rightSection" />
    </NativeSelect>
      </Group>
    
    
    </Card>

    <Title order={2} color="white" align="center" style="padding-top: 20px;">Pending / In Progress Tasks</Title>
    {#each $tasks as task (task.ID)}
      {#if task.Status !== "Completed"}
        <Card withBorder shadow="sm" padding="md" mt="md" class={task.Status === "Completed" ? "inactive" : ""}>
          <Group position="apart">
            <Title order={3}>{task.Title}</Title>
            <Group position="apart">
            <Badge color={task.Priority === "Low" ? "green" : task.Status === "Medium" ? "yellow" : "red"}>
                {task.Priority}
            </Badge>
            <Badge color={task.Status === "Pending" ? "blue" : task.Status === "In progress" ? "green" : "gray"}>
              {task.Status}
            </Badge>
        </Group>
           </Group>
          <p>{task.Description}</p>
          <Group mt="md" position="apart">
            <small>Deadline: {task.Deadline}</small>

            <Group position="apart">
              <Button color="green" on:click={() => completeTask(task)}>
                Done
              </Button>
              <Button color="blue" on:click={() => openUpdatePopUp(task)}>
                Update
              </Button>
              <Button color="red" on:click={() => openDeletePopUp(task)}>
                Delete
              </Button>
            </Group>
          </Group>
        </Card>
      {/if}
    {/each}

    <Title order={2} color="white" align="center" style="padding-top: 20px;"
      >Completed Tasks</Title
    >
    {#each completedTasks as task (task.ID)}
      <Card withBorder shadow="sm" padding="md" mt="md" class="inactive">
        <Group position="apart">
          <Title order={3}>{task.Title}</Title>
          <Badge color="gray">
            {task.Status}
          </Badge>
        </Group>
        <p>{task.Description}</p>
        <Group mt="md" position="apart">
          <small>Deadline: {task.Deadline}</small>

          <Group position="apart">
            <Button color="green" on:click={() => uncompleteTask(task)}>
              Undone
            </Button>
            <Button color="red" on:click={() => openDeletePopUp(task)}>
              Delete
            </Button>
          </Group>
        </Group>
      </Card>
    {/each}

    <Modal
      bind:opened={$showDeletePopUp}
      withCloseButton={false}
      overlayOpacity={0.55}
      overlayBlur={3}
    >
      <div class="modal-content">
        <h3>Are you sure you want to delete this task? 🤔</h3>
        <Group position="apart">
          <Button color="red" on:click={confirmDelete}>Yes, Delete</Button>
          <Button on:click={closeDeletePopUp}>Cancel</Button>
        </Group>
      </div>
    </Modal>

    <Modal
      bind:opened={$showUpdatePopUp}
      withCloseButton={false}
      overlayOpacity={0.55}
      overlayBlur={3}
    >
      <div class="modal-content">
        <h3>Update priority and status of the task</h3>
        <Group position="apart">
          <NativeSelect
            placeholder="Priority"
            bind:value={$taskToUpdate.priority}
            data={["Low", "Medium", "High"]}
          >
            <svelte:component this={ChevronDown} slot="rightSection" />
          </NativeSelect>
          <NativeSelect
            placeholder="Status"
            bind:value={$taskToUpdate.status}
            data={["Pending", "In progress"]}
          >
            <svelte:component this={ChevronDown} slot="rightSection" />
          </NativeSelect>
          <Group position="apart">
            <Button color="blue" on:click={confirmUpdate}>Update</Button>
            <Button color="blue" on:click={closeUpdatePopUp}>Cancel</Button>
          </Group>
        </Group>
      </div>
    </Modal>
  </main>
</SvelteUIProvider>

<style>
  main {
    padding: 10%;
  }
  .inactive {
    opacity: 0.5;
    pointer-events: none;
  }
</style>
