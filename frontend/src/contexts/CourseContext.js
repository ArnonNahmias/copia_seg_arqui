import React, { createContext, useState, useEffect } from 'react';

export const CourseContext = createContext();

export const CourseProvider = ({ children }) => {
  const [courses, setCourses] = useState([]);
  const [enrolledCourses, setEnrolledCourses] = useState([]);

  useEffect(() => {
    const fetchCourses = async () => {
      try {
        const response = await fetch('http://localhost:8080/courses');
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        const data = await response.json();
        setCourses(data);
      } catch (error) {
        console.error('Error fetching courses:', error);
      }
    };

    fetchCourses();
  }, []);

  const enrollCourse = (course) => {
    setEnrolledCourses([...enrolledCourses, course]);
  };

  const unenrollCourse = (course) => {
    setEnrolledCourses(enrolledCourses.filter(c => c.id !== course.id));
  };

  return (
    <CourseContext.Provider value={{ courses, enrolledCourses, enrollCourse, unenrollCourse }}>
      {children}
    </CourseContext.Provider>
  );
};
