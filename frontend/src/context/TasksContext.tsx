import React, { createContext, useState, useMemo } from 'react';
import { useTasks } from '../hooks/useTasks.ts';
import type { Task } from '../services/taskService';

export type TaskFilter = 'all' | 'active' | 'completed';

interface TasksContextProps {
  tasks: Task[];
  loading: boolean;
  filter: TaskFilter;
  setFilter: (filter: TaskFilter) => void;
  addTask: (text: string) => void;
  updateTask: (task: Task) => void;
  deleteTask: (id: string) => void;
}

export const TasksContext = createContext<TasksContextProps>({} as TasksContextProps);

export const TasksProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const { tasks, loading, addTask, updateTask, deleteTask } = useTasks();
  const [filter, setFilter] = useState<TaskFilter>('all');

  const filteredTasks = useMemo(() => {
    if (filter === 'active') return tasks.filter(t => !t.completed);
    if (filter === 'completed') return tasks.filter(t => t.completed);
    return tasks;
  }, [tasks, filter]);

  return (
    <TasksContext.Provider
      value={{
        tasks: filteredTasks,
        loading,
        filter,
        setFilter,
        addTask,
        updateTask,
        deleteTask,
      }}
    >
      {children}
    </TasksContext.Provider>
  );
}; 