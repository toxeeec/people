import { idSchema, paginateSchema } from '@people/common';
import express, { Request } from 'express';
import { FriendRequest } from '../entity/FriendRequest';
import IUser from '../entity/IUser';
import ApiError from '../helpers/ApiError';
import paginate from '../helpers/paginate';
import authenticateToken from '../middlewares/authenticateToken';
import getEntityById from '../middlewares/getEntityById';
import validateParams from '../middlewares/validateParams';

const router = express.Router();

router.post(
  '/:id',
  authenticateToken,
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
        relations: ['sender', 'receiver'],
        where: [
          { sender: user.id, receiver: target.id },
          { sender: target.id, receiver: user.id },
        ],
      });
      if (existingRequest) {
        if (existingRequest.sender.id === user.id) {
          return next(
            new ApiError(400, 'Request to this user was already sent')
          );
        }
        if (existingRequest.receiver.id === user.id) {
          return next(
            new ApiError(400, 'This user already sent u a friend request')
          );
        }
      }
      const request = FriendRequest.create();
      request.sender = { id: user.id } as IUser;
      request.receiver = { id: target.id } as IUser;
      await request.save();
      res.json({ message: 'Request sent successfully' });
    } catch (err) {
      return next(ApiError.internal());
    }
  }
);

router.get(
  '/sent',
  authenticateToken,
  validateParams(paginateSchema),
  async (req: Request, res, next) => {
    const page = req.params.page && parseInt(req.params.page, 10);
    const limit = req.params.limit && parseInt(req.params.limit, 10);
    const result = await paginate(next, FriendRequest, {
      queryOptions: {
        where: { sender: req.user!.id },
      },
      page,
      limit,
    });
    res.json(result);
  }
);

router.get(
  '/received',
  authenticateToken,
  validateParams(paginateSchema),
  async (req: Request, res, next) => {
    const page = req.params.page && parseInt(req.params.page, 10);
    const limit = req.params.limit && parseInt(req.params.limit, 10);
    const result = await paginate(next, FriendRequest, {
      queryOptions: {
        where: { receiver: req.user!.id },
      },
      page,
      limit,
    });
    res.json(result);
  }
);

export default router;
