import React, { useEffect, useState } from 'react';
import { fetchCourses } from '../services/api';
import { Box, Card, CardContent, CardMedia, Typography, Grid, CircularProgress, Alert } from '@mui/material';

const HomePage = () => {
  const [courses, setCourses] = useState([]);
  const [error, setError] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const getCourses = async () => {
      try {
        const coursesData = await fetchCourses();
        setCourses(coursesData);
      } catch (error) {
        console.error('Error fetching courses', error);
        setError('Error fetching courses');
      } finally {
        setLoading(false);
      }
    };

    getCourses();
  }, []);

  if (loading) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>
        <CircularProgress />
      </Box>
    );
  }

  if (error) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>
        <Alert severity="error">{error}</Alert>
      </Box>
    );
  }

  return (
    <Box sx={{ flexGrow: 1, p: 3 }}>
      <Typography variant="h4" component="h1" gutterBottom>
        Available Courses
      </Typography>
      <Grid container spacing={3}>
        {courses.map(course => (
          <Grid item xs={12} sm={6} md={4} key={course.ID}>
            <Card>
              <CardMedia
                component="img"
                height="auto" // esto es para acomodar el tamanio de las imagenes
                image={course.imageURL} // Utiliza la URL de la imagen del curso
                alt={course.nombre}
              />
              <CardContent>
                <Typography gutterBottom variant="h5" component="div">
                  {course.nombre}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  Difficulty: {course.dificultad}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  Price: ${course.precio}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  Address: {course.direccion}
                </Typography>
              </CardContent>
            </Card>
          </Grid>
        ))}
      </Grid>
    </Box>
  );
};

export default HomePage;
