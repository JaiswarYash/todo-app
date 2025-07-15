import React, { useContext, useState } from 'react';
import { TasksContext } from '../context/TasksContext.tsx';
import type { Task } from '../services/taskService.ts';

interface TaskItemProps {
  task: Task;
}

const TaskItem: React.FC<TaskItemProps> = ({ task }) => {
  const { updateTask, deleteTask } = useContext(TasksContext);
  const [isEditing, setIsEditing] = useState(false);
  const [editTitle, setEditTitle] = useState(task.title);

  const handleToggle = () => {
    updateTask({ ...task, completed: !task.completed });
  };

  const handleEdit = () => {
    if (editTitle.trim()) {
      updateTask({ ...task, title: editTitle });
      setIsEditing(false);
    }
  };

  return (
    <div className={`task-item${task.completed ? ' completed' : ''}`}>
      <input type="checkbox" checked={task.completed} onChange={handleToggle} />
      {isEditing ? (
        <input
          type="text"
          value={editTitle}
          onChange={e => setEditTitle(e.target.value)}
          onBlur={handleEdit}
          onKeyDown={e => e.key === 'Enter' && handleEdit()}
          autoFocus
        />
      ) : (
        <span onDoubleClick={() => setIsEditing(true)}>{task.title}</span>
      )}
      <button className="delete-btn" onClick={() => deleteTask(task.id)}>
        Delete
      </button>
    </div>
  );
};

export default TaskItem; 