import IUser from './IUser';

export interface IFriendRequest {
  id: number;
  sender: IUser;
  receiver: IUser;
}
