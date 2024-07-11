// services/api.js
import axios from "axios";

const API_URL = "http://localhost:8080"; // Asegúrate de que esta URL sea correcta

export const fetchCourses = async () => {
  try {
    const response = await axios.get(`${API_URL}/courses`);
    return response.data;
  } catch (error) {
    console.error("Error fetching courses", error);
    throw error;
  }
};

export const loginUser = async (username, password) => {
  try {
    const response = await axios.post(`${API_URL}/login`, {
      username,
      password,
    });
    return response.data;
  } catch (error) {
    console.error("Error logging in", error);
    throw error;
  }
};

export const registerUser = async (username, password) => {
  try {
    const response = await axios.post(`${API_URL}/register`, {
      username,
      password,
    });
    return response.data;
  } catch (error) {
    console.error("Error registering", error);
    throw error;
  }
};

// Nueva función para enviar un comentario
export const addComment = async (courseId, userId, content) => {
  try {
    const response = await axios.post(`${API_URL}/comments`, {
      course_id: courseId,
      user_id: userId,
      content: content,
    });
    return response.data;
  } catch (error) {
    console.error("Error adding comment", error);
    throw error;
  }
};

// Nueva función para obtener comentarios de un curso
export const fetchComments = async (courseId) => {
  try {
    const response = await axios.get(`${API_URL}/courses/${courseId}/comments`);
    return response.data;
  } catch (error) {
    console.error("Error fetching comments", error);
    throw error;
  }
};
