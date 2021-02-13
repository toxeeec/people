import express from 'express';

const router = express.Router();

router.post(':id', async (req, res, next) => {
  res.send({});
});

export default router;
