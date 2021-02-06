import { authenticateUserSchema, registerUserSchema } from '@people/common';
import argon2 from 'argon2';
import express from 'express';
import { User } from '../entity/User';
import ApiError from '../helpers/ApiError';
import createTokens from '../helpers/createTokens';
import validateBody from '../middlewares/validateBody';

const router = express.Router();

router.post(
  '/register',
  validateBody(registerUserSchema),
  async (req, res, next) => {
    try {
      const { name, surname, email, password } = req.body;
      const emailCount = await User.count({ email });
      if (emailCount) return next(new ApiError(400, 'E-mail is already taken'));
      const hashedPassword = await argon2.hash(password);
      const user = User.create({
        name,
        surname,
        email,
        password: hashedPassword,
      });
      await user.save();
      const { accessToken, refreshToken } = await createTokens(user);
      res.status(201).json({ name, surname, email, accessToken, refreshToken });
    } catch (err) {
      next(ApiError.internal());
    }
  }
);

router.post(
  '/login',
  validateBody(authenticateUserSchema),
  async (req, res, next) => {
    try {
      const { email, password } = req.body;
      const user = await User.findOne({ email });
      if (!user) {
        return next(new ApiError(401, 'Wrong username or password'));
      }
      if (!(await argon2.verify(user.password, password))) {
        return next(new ApiError(401, 'Wrong username or password'));
      }
      const { name, surname } = user;
      const { accessToken, refreshToken } = await createTokens(user);
      res.status(200).json({ name, surname, email, accessToken, refreshToken });
    } catch (err) {
      next(ApiError.internal());
    }
  }
);

export default router;
