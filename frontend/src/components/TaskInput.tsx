import React, { useState, useContext } from 'react';
import { TasksContext } from '../context/TasksContext.tsx';

const TaskInput: React.FC = () => {
  const [title, setTitle] = useState('');
  const { addTask } = useContext(TasksContext) as {
    addTask: (title: string) => void;
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (title.trim()) {
      addTask(title);
      setTitle('');
    }
  };

  return (
    <form className="task-input" onSubmit={handleSubmit}>
      <input
        type="text"
        value={title}
        onChange={e => setTitle(e.target.value)}
        placeholder="Add a new task..."
      />
      <button type="submit">Add</button>
    </form>
  );
};

export default TaskInput; 