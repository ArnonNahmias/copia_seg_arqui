// components/Comments.js
import React, { useState, useEffect } from "react";
import {
  Box,
  Typography,
  TextField,
  Button,
  List,
  ListItem,
  ListItemText,
} from "@mui/material";
import { fetchComments, addComment } from "../services/api";

const Comments = ({ courseId, userId }) => {
  const [comments, setComments] = useState([]);
  const [content, setContent] = useState("");
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const getComments = async () => {
      try {
        const commentsData = await fetchComments(courseId);
        setComments(commentsData);
      } catch (error) {
        console.error("Error fetching comments", error);
        setError("Error fetching comments");
      } finally {
        setLoading(false);
      }
    };

    getComments();
  }, [courseId]);

  const handleAddComment = async () => {
    if (content.trim() === "") return;
    try {
      await addComment(courseId, userId, content);
      setComments([
        ...comments,
        { content, user_id: userId, created_at: new Date() },
      ]);
      setContent("");
    } catch (error) {
      console.error("Error adding comment", error);
      setError("Error adding comment");
    }
  };

  if (loading) return <Typography>Loading comments...</Typography>;
  if (error) return <Typography>{error}</Typography>;

  return (
    <Box>
      <Typography variant="h6">Comments</Typography>
      <List>
        {comments.map((comment, index) => (
          <ListItem key={index}>
            <ListItemText
              primary={comment.content}
              secondary={`User ID: ${comment.user_id} - ${new Date(
                comment.created_at
              ).toLocaleString()}`}
            />
          </ListItem>
        ))}
      </List>
      <TextField
        label="Add a comment"
        variant="outlined"
        fullWidth
        margin="normal"
        value={content}
        onChange={(e) => setContent(e.target.value)}
      />
      <Button variant="contained" color="primary" onClick={handleAddComment}>
        Submit
      </Button>
    </Box>
  );
};

export default Comments;
