import { useState, useEffect } from 'react';
import * as taskService from '../services/taskService.ts';
import type { Task } from '../services/taskService.ts';

export const useTasks = () => {
  const [tasks, setTasks] = useState<Task[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    taskService.getTasks().then(data => {
      setTasks(data);
      setLoading(false);
    });
  }, []);

  const addTask = async (title: string) => {
    const newTask = await taskService.addTask(title);
    setTasks(prev => [...prev, newTask]);
  };

  const updateTask = async (task: Task) => {
    const updated = await taskService.updateTask(task);
    setTasks(prev => prev.map(t => (t.id === updated.id ? updated : t)));
  };

  const deleteTask = async (id: string) => {
    await taskService.deleteTask(id);
    setTasks(prev => prev.filter(t => t.id !== id));
  };

  return { tasks, loading, addTask, updateTask, deleteTask };
}; 