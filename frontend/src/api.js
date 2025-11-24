const API_URL = 'http://localhost:8080/api/courses';

export const getCourses = async (author = '') => {
  const url = author ? `${API_URL}?author=${encodeURIComponent(author)}` : API_URL;
  const response = await fetch(url);
  if (!response.ok) throw new Error('Failed to fetch courses');
  return response.json();
};

export const createCourse = async (course) => {
  const response = await fetch(API_URL, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(course),
  });
  if (!response.ok) {
    const error = await response.json();
    throw new Error(error.error || 'Failed to create course');
  }
  return response.json();
};

export const deleteCourse = async (id) => {
  const response = await fetch(`${API_URL}/${id}`, {
    method: 'DELETE',
  });
  if (!response.ok) throw new Error('Failed to delete course');
};

export const updateCourse = async (id, course) => {
  const response = await fetch(`${API_URL}/${id}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(course),
  });
  if (!response.ok) {
    const error = await response.json();
    throw new Error(error.error || 'Failed to update course');
  }
  return response.json();
};
