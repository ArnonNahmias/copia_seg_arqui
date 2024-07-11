import React, { useEffect, useState, useContext } from "react";
import { useParams } from "react-router-dom";
import { Box, Typography, CircularProgress, Alert } from "@mui/material";
import { fetchCourse } from "../services/api";
import Comments from "../components/Comments";
import { UserContext } from "../context/UserContext"; // Importa el UserContext

const CoursePage = () => {
  const { id } = useParams();
  const [course, setCourse] = useState(null);
  const [error, setError] = useState(null);
  const [loading, setLoading] = useState(true);
  const { user } = useContext(UserContext); // Obtén el usuario del contexto

  useEffect(() => {
    const getCourse = async () => {
      try {
        const courseData = await fetchCourse(id);
        setCourse(courseData);
      } catch (error) {
        console.error("Error fetching course", error);
        setError("Error fetching course");
      } finally {
        setLoading(false);
      }
    };

    getCourse();
  }, [id]);

  if (loading) {
    return (
      <Box
        sx={{
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          height: "100vh",
        }}
      >
        <CircularProgress />
      </Box>
    );
  }

  if (error) {
    return (
      <Box
        sx={{
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          height: "100vh",
        }}
      >
        <Alert severity="error">{error}</Alert>
      </Box>
    );
  }

  return (
    <Box sx={{ flexGrow: 1, p: 3 }}>
      <Typography variant="h4" component="h1" gutterBottom>
        {course.nombre}
      </Typography>
      <Typography variant="body1" paragraph>
        Difficulty: {course.dificultad}
      </Typography>
      <Typography variant="body1" paragraph>
        Price: ${course.precio}
      </Typography>
      <Typography variant="body1" paragraph>
        Address: {course.direccion}
      </Typography>
      <Comments courseId={id} userId={user?.id} />
    </Box>
  );
};

export default CoursePage;
