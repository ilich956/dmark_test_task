<script>
import { onMount } from 'svelte';
import {SvelteUIProvider , Button, Input, Textarea, Card, Title, Group, Badge, NativeSelect, Alert, Modal } from '@svelteuidev/core';
import { ChevronDown, InfoCircled } from 'radix-icons-svelte';
import { writable } from 'svelte/store';
import { fetchTasks, addTask, deleteTask, updateTask, tasks, newTask, isAlert } from "./taskService.js";

  const TASK_STATUS = ["Pending", "In progress", "Completed"];
  const TASK_PRIORITY = ["Low", "Medium", "High"];

  let showDeletePopUp = writable(false);
  let showUpdatePopUp = writable(false);

  let taskToDelete = writable(null);
  let taskToUpdate = writable(null);

// sorting order and priority ranking for sorting
  let sortOrder = writable("ascending");
  const priorityOrder = {
        "Low": 1,
        "Medium": 2,
        "High": 3
    };

  // execute when component is initialized
  onMount(async () => {
    fetchTasks();
  });

  // Function to sort tasks based on priority
  function sortTasksByPriority() {
    tasks.update(currentTasks => {
        const order = $sortOrder === "ascending" ? 1 : -1;
        return [...currentTasks].sort((a, b) => (priorityOrder[a.Priority] - priorityOrder[b.Priority]) * order);
    });
}

  // Open delete confirmation modal
  function openDeletePopUp(task) {
    taskToDelete.set(task.ID);
    showDeletePopUp.set(true);
  }

  // Open update modal and set the selected task details
  function openUpdatePopUp(task) {
    taskToUpdate.set({
      ID: task.ID,
      priority: task.Priority,
      status: task.Status,
    });
    showUpdatePopUp.set(true);
  }

  //close modals
  function closeDeletePopUp() {
    showDeletePopUp.set(false);
  }

  function closeUpdatePopUp() {
    showUpdatePopUp.set(false);
  }

  //confirm deletion of task
  async function confirmDelete() {
    const id = $taskToDelete;
    deleteTask(id);
    closeDeletePopUp();
  }

  //Confirm deletion of task
  async function confirmUpdate() {
    const task = $taskToUpdate;
    updateTask(task.ID, task.priority, task.status);
    closeUpdatePopUp();
  }

  //Reactive statement to filter tasks
  $: completedTasks = $tasks.filter((task) => task.Status === "Completed");
  $: console.log($tasks);
</script>

<SvelteUIProvider>
  <main>
     <!-- Alert for missing task title -->
    {#if $isAlert}
      <Alert icon={InfoCircled} title="Oopsie!" variant="filled" radius="xs" withCloseButton on:close={() => isAlert.set(false)}>
        Enter a title for the task!
      </Alert>
    {/if}

    <!-- Main Title -->
    <Title order={1} align="center" style="color: aliceblue; margin-bottom:30px">Simple todo list 📝</Title
    >

    <!-- Task Input Card -->
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
          data={TASK_PRIORITY}
        >
          <svelte:component this={ChevronDown} slot="rightSection" />
        </NativeSelect>
        <NativeSelect
          placeholder="Status"
          bind:value={$newTask.status}
          data={TASK_STATUS}
        >
          <svelte:component this={ChevronDown} slot="rightSection" />
        </NativeSelect>
      </Group>
      <Textarea
        placeholder="Task description (optional)"
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

     <!-- Available Tasks -->
    <Title order={2} color="white" align="center" style="padding-top: 20px;">Available Tasks</Title>
    {#each $tasks as task (task.ID)}
      {#if task.Status !== "Completed"}
        <Card withBorder shadow="sm" padding="md" mt="md" class={task.Status === "Completed" ? "inactive" : ""}>
          <Group position="apart">
            <Title order={3}>{task.Title}</Title>
            <Group position="apart">
            <Badge color={task.Priority === "Low" ? "green" : task.Priority === "Medium" ? "yellow" : "red"}>
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
              <Button color="green" on:click={() => updateTask(task.ID, task.Priority, "Completed")}>
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

    <!-- Completed tasks Section -->
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
            <Button color="green" on:click={() => updateTask(task.ID, task.Priority, "In progress")}>
              Undone
            </Button>
            <Button color="red" on:click={() => openDeletePopUp(task)}>
              Delete
            </Button>
          </Group>
        </Group>
      </Card>
    {/each}

    <!-- Delete confirmation modal -->
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

    <!-- Update task modal -->
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
            data={TASK_PRIORITY}
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
