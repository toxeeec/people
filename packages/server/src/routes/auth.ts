import { registerUserSchema } from '@people/common';
import express from 'express';
import { User } from '../entity/User';
import ApiError from '../helpers/ApiError';
import validateBody from '../middlewares/validateBody';

const router = express.Router();

router.post(
  '/register',
  validateBody(registerUserSchema),
  async (req, res, next) => {
    try {
      const { name, surname, email, password } = req.body;
      const user = User.create({ name, surname, email, password });
      await user.save();
      res.status(201).json(user);
    } catch (err) {
      console.log(err);
      next(ApiError.internal());
    }
  }
);

export default router;
