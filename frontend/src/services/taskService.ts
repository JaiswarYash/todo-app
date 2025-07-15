import axios from 'axios';

const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api';

export interface Task {
  id: string;
  title: string;
  completed: boolean;
  created_at?: string;
  updated_at?: string;
}

export const getTasks = async (): Promise<Task[]> => {
  const res = await axios.get(`${API_URL}/tasks`);
  return res.data;
};

export const addTask = async (title: string): Promise<Task> => {
  const res = await axios.post(`${API_URL}/tasks`, { title });
  return res.data;
};

export const updateTask = async (task: Task): Promise<Task> => {
  const res = await axios.put(`${API_URL}/tasks/${task.id}`, {
    title: task.title,
    completed: task.completed,
  });
  return res.data;
};

export const deleteTask = async (id: string): Promise<void> => {
  await axios.delete(`${API_URL}/tasks/${id}`);
}; 