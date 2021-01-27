import { registerUserSchema } from '@people/common';
import express from 'express';
import validateBody from '../middlewares/validateBody';

const router = express.Router();

router.post(
  '/register',
  validateBody(registerUserSchema),
  async (req, res, next) => {},
);

export default router;
