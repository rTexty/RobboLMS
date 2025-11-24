import React, { useState, useEffect } from 'react';

const CourseForm = ({ onCourseAdded, editingCourse, onCourseUpdated, onCancelEdit }) => {
    const [title, setTitle] = useState('');
    const [description, setDescription] = useState('');
    const [author, setAuthor] = useState('');
    const [error, setError] = useState('');

    useEffect(() => {
        if (editingCourse) {
            setTitle(editingCourse.title);
            setDescription(editingCourse.description);
            setAuthor(editingCourse.author || '');
        } else {
            setTitle('');
            setDescription('');
            setAuthor('');
        }
    }, [editingCourse]);

    const handleSubmit = async (e) => {
        e.preventDefault();
        setError('');

        if (title.length < 3) {
            setError('Title must be at least 3 characters long');
            return;
        }

        try {
            if (editingCourse) {
                await onCourseUpdated(editingCourse.id, { title, description, author });
            } else {
                await onCourseAdded({ title, description, author });
            }
            setTitle('');
            setDescription('');
            setAuthor('');
        } catch (err) {
            setError(err.message);
        }
    };

    return (
        <form onSubmit={handleSubmit} className="course-form">
            <h3>{editingCourse ? 'Edit Course' : 'Add New Course'}</h3>
            {error && <div className="error">{error}</div>}
            <div className="form-group">
                <label>Title (min 3 chars):</label>
                <input
                    type="text"
                    value={title}
                    onChange={(e) => setTitle(e.target.value)}
                    required
                />
            </div>
            <div className="form-group">
                <label>Description:</label>
                <textarea
                    value={description}
                    onChange={(e) => setDescription(e.target.value)}
                />
            </div>
            <div className="form-group">
                <label>Author:</label>
                <input
                    type="text"
                    value={author}
                    onChange={(e) => setAuthor(e.target.value)}
                />
            </div>
            <button type="submit">{editingCourse ? 'Update Course' : 'Add Course'}</button>
            {editingCourse && (
                <button type="button" onClick={onCancelEdit} className="cancel-btn">
                    Cancel
                </button>
            )}
        </form>
    );
};

export default CourseForm;
