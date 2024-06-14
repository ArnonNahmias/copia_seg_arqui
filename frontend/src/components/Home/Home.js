import React, { useContext, useEffect, useState } from 'react';
import { Container, Row, Col, Card, Button, Alert, Form } from 'react-bootstrap';
import { CourseContext } from '../../contexts/CourseContext';
import { useNavigate } from 'react-router-dom';
import './Home.scss';

const Home = ({ userRole }) => {
  const { courses, enrollCourse, enrolledCourses } = useContext(CourseContext);
  const [message, setMessage] = useState(null);
  const [searchTerm, setSearchTerm] = useState('');
  const [filteredCourses, setFilteredCourses] = useState([]);
  const navigate = useNavigate();

  useEffect(() => {
    if (courses) {
      setFilteredCourses(courses.filter(course =>
        (course.nombre && course.nombre.toLowerCase().includes(searchTerm.toLowerCase())) ||
        (course.direccion && course.direccion.toLowerCase().includes(searchTerm.toLowerCase()))
      ));
    }
  }, [searchTerm, courses]);

  const handleEnroll = async (course) => {
    if (userRole !== 'commonUser') {
      setMessage('You must be logged in as a common user to subscribe to courses.');
      setTimeout(() => {
        navigate('/login');
      }, 1500);
      return;
    }

    if (enrolledCourses.some(c => c.nombre === course.nombre)) {
      setMessage(`You are already enrolled in the course: ${course.nombre}`);
    } else {
      try {
        const userId = localStorage.getItem('userId');
        const response = await fetch('http://localhost:8080/subscriptions', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${localStorage.getItem('token')}`
          },
          body: JSON.stringify({ user_id: userId, course_id: course.id })
        });
        if (!response.ok) {
          throw new Error('Failed to enroll');
        }
        enrollCourse(course);
        setMessage(`Congratulations on enrolling in the course: ${course.nombre}`);
      } catch (error) {
        setMessage({ type: 'danger', text: error.message });
      }
    }
    setTimeout(() => setMessage(null), 3000);
  };

  const handleSearchChange = (e) => {
    setSearchTerm(e.target.value);
  };

  return (
    <Container className="common-user-view">
      {message && <Alert variant={message.type === 'danger' ? 'danger' : 'success'}>{message}</Alert>}
      <Form.Control
        type="text"
        placeholder="Search courses"
        value={searchTerm}
        onChange={handleSearchChange}
        className="mb-4"
      />
      <Row>
        {filteredCourses.length > 0 ? (
          filteredCourses.map((course, index) => (
            <Col key={index} xs={12} md={6} lg={4} className="course-col">
              <Card className="course-card">
                {course.imageURL && <Card.Img variant="top" src={course.imageURL} alt={course.nombre} />}
                <Card.Body>
                  <Card.Title>{course.nombre}</Card.Title>
                  <Card.Text><strong>Dificultad:</strong> {course.dificultad}</Card.Text>
                  <Card.Text><strong>Precio:</strong> ${course.precio}</Card.Text>
                  <Card.Text><strong>Direccion:</strong> {course.direccion}</Card.Text>
                  <Button variant="primary" onClick={() => handleEnroll(course)}>Enroll</Button>
                </Card.Body>
              </Card>
            </Col>
          ))
        ) : (
          <p>No courses available</p>
        )}
      </Row>
    </Container>
  );
};

export default Home;
