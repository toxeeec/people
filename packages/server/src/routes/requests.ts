import { idSchema, paginateSchema } from '@people/common';
import express, { Request } from 'express';
import { FriendRequest, FriendRequestStatus } from '../entity/FriendRequest';
import ApiError from '../helpers/ApiError';
import paginate from '../helpers/paginate';
import authenticateToken from '../middlewares/authenticateToken';
import getEntityById from '../middlewares/getEntityById';
import validateParams from '../middlewares/validateParams';

const router = express.Router();

router.use(authenticateToken);

router.get(
  '/sent',
  validateParams(paginateSchema),
  async (req: Request, res, next) => {
    const page = req.params.page && parseInt(req.params.page, 10);
    const limit = req.params.limit && parseInt(req.params.limit, 10);
    const result = await paginate(next, FriendRequest, {
      queryOptions: {
        where: { senderId: req.user!.id },
      },
      page,
      limit,
    });
    res.json(result);
  }
);

router.get(
  '/received',
  validateParams(paginateSchema),
  async (req: Request, res, next) => {
    const page = req.params.page && parseInt(req.params.page, 10);
    const limit = req.params.limit && parseInt(req.params.limit, 10);
    const result = await paginate(next, FriendRequest, {
      queryOptions: {
        where: { receiverId: req.user!.id },
      },
      page,
      limit,
    });
    res.json(result);
  }
);

router.post(
  '/:id',
  validateParams(idSchema),
  getEntityById('User'),
  async (req: Request, res, next) => {
    const target = req.entity;
    const { user } = req;
    if (target.id === user.id) {
      return next(
        new ApiError(400, 'You cannot send a friend request to yourself')
      );
    }
    try {
      const existingRequest = await FriendRequest.findOne({
        where: [
          { senderId: user.id, receiverId: target.id },
          { senderId: target.id, receiverId: user.id },
        ],
      });
      if (existingRequest) {
        if (existingRequest.status === FriendRequestStatus.ACCEPTED)
          return next(
            new ApiError(400, 'You are already friends with this user')
          );
        if (existingRequest.senderId === user.id) {
          return next(
            new ApiError(400, 'Request to this user was already sent')
          );
        }
        if (existingRequest.receiverId === user.id) {
          return next(
            new ApiError(400, 'This user already sent u a friend request')
          );
        }
      }
      await FriendRequest.create({
        senderId: user.id,
        receiverId: target.id,
      }).save();
      res.status(201).json({ message: 'Request sent successfully' });
    } catch (err) {
      return next(ApiError.internal());
    }
  }
);

router.delete(
  '/:id',
  validateParams(idSchema),
  async (req: Request, res, next) => {
    const { user } = req;
    try {
      await FriendRequest.delete({
        senderId: user.id,
        receiverId: req.params.id,
      });
      res.json({ message: 'Request deleted successfully' });
    } catch (err) {
      return next(ApiError.internal());
    }
  }
);

router.patch(
  '/:id',
  validateParams(idSchema),
  getEntityById('User'),
  async (req, res, next) => {
    res.json({});
  }
);

export default router;
