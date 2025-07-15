import React, { useContext } from 'react';
import { TasksProvider, TasksContext } from './context/TasksContext.tsx';
import TaskInput from './components/TaskInput.tsx';
import FilterBar from './components/FilterBar.tsx';
import TaskItem from './components/TaskItem.tsx';
import './styles/index.css';

const TaskList: React.FC = () => {
  const { tasks } = useContext(TasksContext);
  return (
    <>
      {(tasks ?? []).map((task: any) => (
        <TaskItem key={task.id} task={task} />
      ))}
    </>
  );
};

const App: React.FC = () => (
  <TasksProvider>
    <div className="app-container">
      <h1>To-Do List</h1>
      <TaskInput />
      <FilterBar />
      <TaskList />
    </div>
  </TasksProvider>
);

export default App;
