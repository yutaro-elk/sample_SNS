import actionCreatorFactory from 'typescript-fsa';
import { Dispatch } from 'redux';

import UserAPI from '../../requests/user';
import { ErrorResponse } from '../../types/errorResponse';
import { TweetType } from '../../types/tweet';
import { UserType } from '../../types/user';
import { ActionTypes } from '../actionTypes';

const actionCreator = actionCreatorFactory();

const userAction = {
  getUser: actionCreator.async<{}, UserType, Error>(ActionTypes.getUser),
  getUserTweets: actionCreator.async<{}, TweetType[], Error>(
    ActionTypes.getUserTweets
  ),
};

export const fetchUser = (uid: string) => async (dispatch: Dispatch) => {
  dispatch(userAction.getUser.started({ params: {} }));
  const userAPI = new UserAPI();
  try {
    const res = await userAPI.getUser(uid);
    if (res.ok) {
      const result = (await res.json()) as UserType;
      dispatch(userAction.getUser.done({ result, params: {} }));
    } else {
      const result = (await res.json()) as ErrorResponse;
      throw new Error(result.massage);
    }
  } catch (err) {
    const error = err as Error;
    dispatch(userAction.getUser.failed({ error, params: {} }));
  }
};

export const fetchUserTweets = (uid: string) => async (dispatch: Dispatch) => {
  dispatch(userAction.getUserTweets.started({ params: {} }));
  const userAPI = new UserAPI();
  try {
    const res = await userAPI.getUserTweets(uid);
    if (res.ok) {
      const response = (await res.json()) as TweetType[];
      const result = response === null ? [] : response;
      dispatch(userAction.getUserTweets.done({ result, params: {} }));
    } else {
      const result = (await res.json()) as ErrorResponse;
      throw new Error(result.massage);
    }
  } catch (err) {
    const error = err as Error;
    dispatch(userAction.getUserTweets.failed({ error, params: {} }));
  }
};

export default userAction;
