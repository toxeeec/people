import { idSchema } from '@people/common';
import express, { Request } from 'express';
import { FriendRequest } from '../entity/FriendRequest';
import IUser from '../entity/IUser';
import ApiError from '../helpers/ApiError';
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
      console.log(err);
      return next(ApiError.internal());
    }
  }
);

export default router;
