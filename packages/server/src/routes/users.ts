import express from 'express';
import { User } from '../entity/User';
import ApiError from '../helpers/ApiError';
import validateId from '../middlewares/validateId';

const router = express.Router();

router.get('/:id', validateId, async (req, res, next) => {
  try {
    const user = await User.findOne({ id: req.params.id });
    if (!user) return next(ApiError.notFound());
    const { id, name, surname } = user;
    return res.json({ id, name, surname });
  } catch (err) {
    next(ApiError.internal());
  }
});

export default router;
