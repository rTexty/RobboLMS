import React, { useState, useEffect } from 'react';
import { getCourses, createCourse, deleteCourse, updateCourse } from './api';
import CourseList from './components/CourseList';
import CourseForm from './components/CourseForm';
import Modal from './components/Modal';
import './index.css';

function App() {
  const [courses, setCourses] = useState([]);
  const [editingCourse, setEditingCourse] = useState(null);
  const [error, setError] = useState('');
  const [authorFilter, setAuthorFilter] = useState('');

  const fetchCourses = async () => {
    try {
      const data = await getCourses(authorFilter);
      setCourses(data || []);
    } catch (err) {
      setError('Failed to load courses');
    }
  };

  useEffect(() => {
    // Debounce the filter to avoid too many requests
    const timer = setTimeout(() => {
      fetchCourses();
    }, 300);
    return () => clearTimeout(timer);
  }, [authorFilter]);

  const handleAddCourse = async (course) => {
    await createCourse(course);
    fetchCourses();
  };

  const handleDeleteCourse = async (id) => {
    await deleteCourse(id);
    fetchCourses();
  };

  const handleUpdateCourse = async (id, course) => {
    await updateCourse(id, course);
    setEditingCourse(null);
    fetchCourses();
  };

  return (
    <div className="app-container">
      <header>
        <h1>Course Management System</h1>
      </header>
      <main>
        {error && <div className="error-banner">{error}</div>}
        <div className="content">
          <div className="left-panel">
            <CourseForm
              onCourseAdded={handleAddCourse}
            />
          </div>
          <div className="right-panel">
            <div className="list-header" style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '15px' }}>
              <h2>Courses</h2>
              <div className="filter-controls">
                <input
                  type="text"
                  placeholder="Filter by Author..."
                  value={authorFilter}
                  onChange={(e) => setAuthorFilter(e.target.value)}
                  style={{ padding: '8px', borderRadius: '4px', border: '1px solid #ddd' }}
                />
              </div>
            </div>
            <CourseList
              courses={courses}
              onDelete={handleDeleteCourse}
              onEdit={setEditingCourse}
            />
          </div>
        </div>
      </main>

      <Modal isOpen={!!editingCourse} onClose={() => setEditingCourse(null)}>
        <CourseForm
          editingCourse={editingCourse}
          onCourseUpdated={handleUpdateCourse}
          onCancelEdit={() => setEditingCourse(null)}
        />
      </Modal>
    </div>
  );
}

export default App;
