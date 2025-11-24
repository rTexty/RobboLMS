import React from 'react';

const CourseList = ({ courses, onDelete, onEdit }) => {
    return (
        <div className="course-list">
            {courses?.map((course) => (
                <div key={course.id} className="course-card">
                    <h4>{course.title}</h4>
                    <p>{course.description}</p>
                    {course.author && <small>Author: {course.author}</small>}
                    <div className="actions">
                        <button onClick={() => onEdit(course)}>Edit</button>
                        <button
                            onClick={() => {
                                if (window.confirm('Are you sure you want to delete this course?')) {
                                    onDelete(course.id);
                                }
                            }}
                            className="delete-btn"
                        >
                            Delete
                        </button>
                    </div>
                </div>
            ))}
        </div>
    );
};

export default CourseList;
