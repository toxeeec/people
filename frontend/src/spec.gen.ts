/**
 * Generated by orval v6.10.3 🍺
 * Do not edit manually.
 * People API
 * OpenAPI spec version: 1.0.0
 */
import {
  useQuery,
  useMutation
} from '@tanstack/react-query'
import type {
  UseQueryOptions,
  UseMutationOptions,
  QueryFunction,
  MutationFunction,
  UseQueryResult,
  QueryKey
} from '@tanstack/react-query'
import type {
  AuthResponse,
  BadRequestResponse,
  AuthUserBodyBody,
  UnauthorizedResponse,
  Tokens,
  ForbiddenResponse,
  TokensBodyBody,
  Posts,
  GetMeFeedParams,
  NoContentResponse,
  NotFoundResponse,
  Error,
  Users,
  GetMeFollowingParams,
  GetMeFollowersParams,
  GetUsersHandleFollowingParams,
  GetUsersHandleFollowersParams,
  Post,
  PostBodyBody,
  GetUsersHandlePostsParams,
  GetPostsPostIDRepliesParams,
  Likes
} from './models'
import { customInstance } from './custom-instance'
import type { ErrorType } from './custom-instance'



// eslint-disable-next-line
  type SecondParameter<T extends (...args: any) => any> = T extends (
  config: any,
  args: infer P,
) => any
  ? P
  : never;

export const postRegister = (
    authUserBodyBody: AuthUserBodyBody,
 options?: SecondParameter<typeof customInstance>,) => {
      return customInstance<AuthResponse>(
      {url: `/register`, method: 'post',
      headers: {'Content-Type': 'application/json', },
      data: authUserBodyBody
    },
      options);
    }
  


    export type PostRegisterMutationResult = NonNullable<Awaited<ReturnType<typeof postRegister>>>
    export type PostRegisterMutationBody = AuthUserBodyBody
    export type PostRegisterMutationError = ErrorType<BadRequestResponse>

    export const usePostRegister = <TError = ErrorType<BadRequestResponse>,
    
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof postRegister>>, TError,{data: AuthUserBodyBody}, TContext>, request?: SecondParameter<typeof customInstance>}
) => {
      const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof postRegister>>, {data: AuthUserBodyBody}> = (props) => {
          const {data} = props ?? {};

          return  postRegister(data,requestOptions)
        }

      return useMutation<Awaited<ReturnType<typeof postRegister>>, TError, {data: AuthUserBodyBody}, TContext>(mutationFn, mutationOptions)
    }
    
export const postLogin = (
    authUserBodyBody: AuthUserBodyBody,
 options?: SecondParameter<typeof customInstance>,) => {
      return customInstance<AuthResponse>(
      {url: `/login`, method: 'post',
      headers: {'Content-Type': 'application/json', },
      data: authUserBodyBody
    },
      options);
    }
  


    export type PostLoginMutationResult = NonNullable<Awaited<ReturnType<typeof postLogin>>>
    export type PostLoginMutationBody = AuthUserBodyBody
    export type PostLoginMutationError = ErrorType<BadRequestResponse | UnauthorizedResponse>

    export const usePostLogin = <TError = ErrorType<BadRequestResponse | UnauthorizedResponse>,
    
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof postLogin>>, TError,{data: AuthUserBodyBody}, TContext>, request?: SecondParameter<typeof customInstance>}
) => {
      const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof postLogin>>, {data: AuthUserBodyBody}> = (props) => {
          const {data} = props ?? {};

          return  postLogin(data,requestOptions)
        }

      return useMutation<Awaited<ReturnType<typeof postLogin>>, TError, {data: AuthUserBodyBody}, TContext>(mutationFn, mutationOptions)
    }
    
export const postRefresh = (
    tokensBodyBody: TokensBodyBody,
 options?: SecondParameter<typeof customInstance>,) => {
      return customInstance<Tokens>(
      {url: `/refresh`, method: 'post',
      headers: {'Content-Type': 'application/json', },
      data: tokensBodyBody
    },
      options);
    }
  


    export type PostRefreshMutationResult = NonNullable<Awaited<ReturnType<typeof postRefresh>>>
    export type PostRefreshMutationBody = TokensBodyBody
    export type PostRefreshMutationError = ErrorType<BadRequestResponse | ForbiddenResponse>

    export const usePostRefresh = <TError = ErrorType<BadRequestResponse | ForbiddenResponse>,
    
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof postRefresh>>, TError,{data: TokensBodyBody}, TContext>, request?: SecondParameter<typeof customInstance>}
) => {
      const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof postRefresh>>, {data: TokensBodyBody}> = (props) => {
          const {data} = props ?? {};

          return  postRefresh(data,requestOptions)
        }

      return useMutation<Awaited<ReturnType<typeof postRefresh>>, TError, {data: TokensBodyBody}, TContext>(mutationFn, mutationOptions)
    }
    
export const getMeFeed = (
    params?: GetMeFeedParams,
 options?: SecondParameter<typeof customInstance>,signal?: AbortSignal
) => {
      return customInstance<Posts>(
      {url: `/me/feed`, method: 'get',
        params, signal
    },
      options);
    }
  

export const getGetMeFeedQueryKey = (params?: GetMeFeedParams,) => [`/me/feed`, ...(params ? [params]: [])];

    
export type GetMeFeedQueryResult = NonNullable<Awaited<ReturnType<typeof getMeFeed>>>
export type GetMeFeedQueryError = ErrorType<UnauthorizedResponse | ForbiddenResponse>

export const useGetMeFeed = <TData = Awaited<ReturnType<typeof getMeFeed>>, TError = ErrorType<UnauthorizedResponse | ForbiddenResponse>>(
 params?: GetMeFeedParams, options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getMeFeed>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey = queryOptions?.queryKey ?? getGetMeFeedQueryKey(params);

  

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getMeFeed>>> = ({ signal }) => getMeFeed(params, requestOptions, signal);

  const query = useQuery<Awaited<ReturnType<typeof getMeFeed>>, TError, TData>(queryKey, queryFn, queryOptions) as  UseQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryKey;

  return query;
}


export const getMeFollowingHandle = (
    handle: string,
 options?: SecondParameter<typeof customInstance>,signal?: AbortSignal
) => {
      return customInstance<NoContentResponse>(
      {url: `/me/following/${handle}`, method: 'get', signal
    },
      options);
    }
  

export const getGetMeFollowingHandleQueryKey = (handle: string,) => [`/me/following/${handle}`];

    
export type GetMeFollowingHandleQueryResult = NonNullable<Awaited<ReturnType<typeof getMeFollowingHandle>>>
export type GetMeFollowingHandleQueryError = ErrorType<BadRequestResponse | UnauthorizedResponse | ForbiddenResponse | NotFoundResponse>

export const useGetMeFollowingHandle = <TData = Awaited<ReturnType<typeof getMeFollowingHandle>>, TError = ErrorType<BadRequestResponse | UnauthorizedResponse | ForbiddenResponse | NotFoundResponse>>(
 handle: string, options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getMeFollowingHandle>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey = queryOptions?.queryKey ?? getGetMeFollowingHandleQueryKey(handle);

  

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getMeFollowingHandle>>> = ({ signal }) => getMeFollowingHandle(handle, requestOptions, signal);

  const query = useQuery<Awaited<ReturnType<typeof getMeFollowingHandle>>, TError, TData>(queryKey, queryFn, {enabled: !!(handle), ...queryOptions}) as  UseQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryKey;

  return query;
}


export const putMeFollowingHandle = (
    handle: string,
 options?: SecondParameter<typeof customInstance>,) => {
      return customInstance<NoContentResponse>(
      {url: `/me/following/${handle}`, method: 'put'
    },
      options);
    }
  


    export type PutMeFollowingHandleMutationResult = NonNullable<Awaited<ReturnType<typeof putMeFollowingHandle>>>
    
    export type PutMeFollowingHandleMutationError = ErrorType<BadRequestResponse | UnauthorizedResponse | ForbiddenResponse | NotFoundResponse | Error>

    export const usePutMeFollowingHandle = <TError = ErrorType<BadRequestResponse | UnauthorizedResponse | ForbiddenResponse | NotFoundResponse | Error>,
    
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof putMeFollowingHandle>>, TError,{handle: string}, TContext>, request?: SecondParameter<typeof customInstance>}
) => {
      const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof putMeFollowingHandle>>, {handle: string}> = (props) => {
          const {handle} = props ?? {};

          return  putMeFollowingHandle(handle,requestOptions)
        }

      return useMutation<Awaited<ReturnType<typeof putMeFollowingHandle>>, TError, {handle: string}, TContext>(mutationFn, mutationOptions)
    }
    
export const deleteMeFollowingHandle = (
    handle: string,
 options?: SecondParameter<typeof customInstance>,) => {
      return customInstance<NoContentResponse>(
      {url: `/me/following/${handle}`, method: 'delete'
    },
      options);
    }
  


    export type DeleteMeFollowingHandleMutationResult = NonNullable<Awaited<ReturnType<typeof deleteMeFollowingHandle>>>
    
    export type DeleteMeFollowingHandleMutationError = ErrorType<BadRequestResponse | UnauthorizedResponse | ForbiddenResponse | NotFoundResponse>

    export const useDeleteMeFollowingHandle = <TError = ErrorType<BadRequestResponse | UnauthorizedResponse | ForbiddenResponse | NotFoundResponse>,
    
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof deleteMeFollowingHandle>>, TError,{handle: string}, TContext>, request?: SecondParameter<typeof customInstance>}
) => {
      const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof deleteMeFollowingHandle>>, {handle: string}> = (props) => {
          const {handle} = props ?? {};

          return  deleteMeFollowingHandle(handle,requestOptions)
        }

      return useMutation<Awaited<ReturnType<typeof deleteMeFollowingHandle>>, TError, {handle: string}, TContext>(mutationFn, mutationOptions)
    }
    
export const getMeFollowing = (
    params?: GetMeFollowingParams,
 options?: SecondParameter<typeof customInstance>,signal?: AbortSignal
) => {
      return customInstance<Users>(
      {url: `/me/following`, method: 'get',
        params, signal
    },
      options);
    }
  

export const getGetMeFollowingQueryKey = (params?: GetMeFollowingParams,) => [`/me/following`, ...(params ? [params]: [])];

    
export type GetMeFollowingQueryResult = NonNullable<Awaited<ReturnType<typeof getMeFollowing>>>
export type GetMeFollowingQueryError = ErrorType<UnauthorizedResponse | ForbiddenResponse>

export const useGetMeFollowing = <TData = Awaited<ReturnType<typeof getMeFollowing>>, TError = ErrorType<UnauthorizedResponse | ForbiddenResponse>>(
 params?: GetMeFollowingParams, options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getMeFollowing>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey = queryOptions?.queryKey ?? getGetMeFollowingQueryKey(params);

  

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getMeFollowing>>> = ({ signal }) => getMeFollowing(params, requestOptions, signal);

  const query = useQuery<Awaited<ReturnType<typeof getMeFollowing>>, TError, TData>(queryKey, queryFn, queryOptions) as  UseQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryKey;

  return query;
}


export const getMeFollowersHandle = (
    handle: string,
 options?: SecondParameter<typeof customInstance>,signal?: AbortSignal
) => {
      return customInstance<NoContentResponse>(
      {url: `/me/followers/${handle}`, method: 'get', signal
    },
      options);
    }
  

export const getGetMeFollowersHandleQueryKey = (handle: string,) => [`/me/followers/${handle}`];

    
export type GetMeFollowersHandleQueryResult = NonNullable<Awaited<ReturnType<typeof getMeFollowersHandle>>>
export type GetMeFollowersHandleQueryError = ErrorType<BadRequestResponse | UnauthorizedResponse | ForbiddenResponse | NotFoundResponse>

export const useGetMeFollowersHandle = <TData = Awaited<ReturnType<typeof getMeFollowersHandle>>, TError = ErrorType<BadRequestResponse | UnauthorizedResponse | ForbiddenResponse | NotFoundResponse>>(
 handle: string, options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getMeFollowersHandle>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey = queryOptions?.queryKey ?? getGetMeFollowersHandleQueryKey(handle);

  

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getMeFollowersHandle>>> = ({ signal }) => getMeFollowersHandle(handle, requestOptions, signal);

  const query = useQuery<Awaited<ReturnType<typeof getMeFollowersHandle>>, TError, TData>(queryKey, queryFn, {enabled: !!(handle), ...queryOptions}) as  UseQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryKey;

  return query;
}


export const getMeFollowers = (
    params?: GetMeFollowersParams,
 options?: SecondParameter<typeof customInstance>,signal?: AbortSignal
) => {
      return customInstance<Users>(
      {url: `/me/followers`, method: 'get',
        params, signal
    },
      options);
    }
  

export const getGetMeFollowersQueryKey = (params?: GetMeFollowersParams,) => [`/me/followers`, ...(params ? [params]: [])];

    
export type GetMeFollowersQueryResult = NonNullable<Awaited<ReturnType<typeof getMeFollowers>>>
export type GetMeFollowersQueryError = ErrorType<UnauthorizedResponse | ForbiddenResponse>

export const useGetMeFollowers = <TData = Awaited<ReturnType<typeof getMeFollowers>>, TError = ErrorType<UnauthorizedResponse | ForbiddenResponse>>(
 params?: GetMeFollowersParams, options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getMeFollowers>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey = queryOptions?.queryKey ?? getGetMeFollowersQueryKey(params);

  

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getMeFollowers>>> = ({ signal }) => getMeFollowers(params, requestOptions, signal);

  const query = useQuery<Awaited<ReturnType<typeof getMeFollowers>>, TError, TData>(queryKey, queryFn, queryOptions) as  UseQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryKey;

  return query;
}


export const getUsersHandleFollowing = (
    handle: string,
    params?: GetUsersHandleFollowingParams,
 options?: SecondParameter<typeof customInstance>,signal?: AbortSignal
) => {
      return customInstance<Users>(
      {url: `/users/${handle}/following`, method: 'get',
        params, signal
    },
      options);
    }
  

export const getGetUsersHandleFollowingQueryKey = (handle: string,
    params?: GetUsersHandleFollowingParams,) => [`/users/${handle}/following`, ...(params ? [params]: [])];

    
export type GetUsersHandleFollowingQueryResult = NonNullable<Awaited<ReturnType<typeof getUsersHandleFollowing>>>
export type GetUsersHandleFollowingQueryError = ErrorType<BadRequestResponse | NotFoundResponse>

export const useGetUsersHandleFollowing = <TData = Awaited<ReturnType<typeof getUsersHandleFollowing>>, TError = ErrorType<BadRequestResponse | NotFoundResponse>>(
 handle: string,
    params?: GetUsersHandleFollowingParams, options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getUsersHandleFollowing>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey = queryOptions?.queryKey ?? getGetUsersHandleFollowingQueryKey(handle,params);

  

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getUsersHandleFollowing>>> = ({ signal }) => getUsersHandleFollowing(handle,params, requestOptions, signal);

  const query = useQuery<Awaited<ReturnType<typeof getUsersHandleFollowing>>, TError, TData>(queryKey, queryFn, {enabled: !!(handle), ...queryOptions}) as  UseQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryKey;

  return query;
}


export const getUsersHandleFollowers = (
    handle: string,
    params?: GetUsersHandleFollowersParams,
 options?: SecondParameter<typeof customInstance>,signal?: AbortSignal
) => {
      return customInstance<Users>(
      {url: `/users/${handle}/followers`, method: 'get',
        params, signal
    },
      options);
    }
  

export const getGetUsersHandleFollowersQueryKey = (handle: string,
    params?: GetUsersHandleFollowersParams,) => [`/users/${handle}/followers`, ...(params ? [params]: [])];

    
export type GetUsersHandleFollowersQueryResult = NonNullable<Awaited<ReturnType<typeof getUsersHandleFollowers>>>
export type GetUsersHandleFollowersQueryError = ErrorType<BadRequestResponse | NotFoundResponse>

export const useGetUsersHandleFollowers = <TData = Awaited<ReturnType<typeof getUsersHandleFollowers>>, TError = ErrorType<BadRequestResponse | NotFoundResponse>>(
 handle: string,
    params?: GetUsersHandleFollowersParams, options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getUsersHandleFollowers>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey = queryOptions?.queryKey ?? getGetUsersHandleFollowersQueryKey(handle,params);

  

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getUsersHandleFollowers>>> = ({ signal }) => getUsersHandleFollowers(handle,params, requestOptions, signal);

  const query = useQuery<Awaited<ReturnType<typeof getUsersHandleFollowers>>, TError, TData>(queryKey, queryFn, {enabled: !!(handle), ...queryOptions}) as  UseQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryKey;

  return query;
}


export const postPosts = (
    postBodyBody: PostBodyBody,
 options?: SecondParameter<typeof customInstance>,) => {
      return customInstance<Post>(
      {url: `/posts`, method: 'post',
      headers: {'Content-Type': 'application/json', },
      data: postBodyBody
    },
      options);
    }
  


    export type PostPostsMutationResult = NonNullable<Awaited<ReturnType<typeof postPosts>>>
    export type PostPostsMutationBody = PostBodyBody
    export type PostPostsMutationError = ErrorType<BadRequestResponse | UnauthorizedResponse | ForbiddenResponse>

    export const usePostPosts = <TError = ErrorType<BadRequestResponse | UnauthorizedResponse | ForbiddenResponse>,
    
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof postPosts>>, TError,{data: PostBodyBody}, TContext>, request?: SecondParameter<typeof customInstance>}
) => {
      const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof postPosts>>, {data: PostBodyBody}> = (props) => {
          const {data} = props ?? {};

          return  postPosts(data,requestOptions)
        }

      return useMutation<Awaited<ReturnType<typeof postPosts>>, TError, {data: PostBodyBody}, TContext>(mutationFn, mutationOptions)
    }
    
export const getPostsPostID = (
    postID: number,
 options?: SecondParameter<typeof customInstance>,signal?: AbortSignal
) => {
      return customInstance<Post>(
      {url: `/posts/${postID}`, method: 'get', signal
    },
      options);
    }
  

export const getGetPostsPostIDQueryKey = (postID: number,) => [`/posts/${postID}`];

    
export type GetPostsPostIDQueryResult = NonNullable<Awaited<ReturnType<typeof getPostsPostID>>>
export type GetPostsPostIDQueryError = ErrorType<BadRequestResponse | NotFoundResponse>

export const useGetPostsPostID = <TData = Awaited<ReturnType<typeof getPostsPostID>>, TError = ErrorType<BadRequestResponse | NotFoundResponse>>(
 postID: number, options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getPostsPostID>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey = queryOptions?.queryKey ?? getGetPostsPostIDQueryKey(postID);

  

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getPostsPostID>>> = ({ signal }) => getPostsPostID(postID, requestOptions, signal);

  const query = useQuery<Awaited<ReturnType<typeof getPostsPostID>>, TError, TData>(queryKey, queryFn, {enabled: !!(postID), ...queryOptions}) as  UseQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryKey;

  return query;
}


export const deletePostsPostID = (
    postID: number,
 options?: SecondParameter<typeof customInstance>,) => {
      return customInstance<NoContentResponse>(
      {url: `/posts/${postID}`, method: 'delete'
    },
      options);
    }
  


    export type DeletePostsPostIDMutationResult = NonNullable<Awaited<ReturnType<typeof deletePostsPostID>>>
    
    export type DeletePostsPostIDMutationError = ErrorType<BadRequestResponse | UnauthorizedResponse | ForbiddenResponse | NotFoundResponse>

    export const useDeletePostsPostID = <TError = ErrorType<BadRequestResponse | UnauthorizedResponse | ForbiddenResponse | NotFoundResponse>,
    
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof deletePostsPostID>>, TError,{postID: number}, TContext>, request?: SecondParameter<typeof customInstance>}
) => {
      const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof deletePostsPostID>>, {postID: number}> = (props) => {
          const {postID} = props ?? {};

          return  deletePostsPostID(postID,requestOptions)
        }

      return useMutation<Awaited<ReturnType<typeof deletePostsPostID>>, TError, {postID: number}, TContext>(mutationFn, mutationOptions)
    }
    
export const getUsersHandlePosts = (
    handle: string,
    params?: GetUsersHandlePostsParams,
 options?: SecondParameter<typeof customInstance>,signal?: AbortSignal
) => {
      return customInstance<Posts>(
      {url: `/users/${handle}/posts`, method: 'get',
        params, signal
    },
      options);
    }
  

export const getGetUsersHandlePostsQueryKey = (handle: string,
    params?: GetUsersHandlePostsParams,) => [`/users/${handle}/posts`, ...(params ? [params]: [])];

    
export type GetUsersHandlePostsQueryResult = NonNullable<Awaited<ReturnType<typeof getUsersHandlePosts>>>
export type GetUsersHandlePostsQueryError = ErrorType<unknown>

export const useGetUsersHandlePosts = <TData = Awaited<ReturnType<typeof getUsersHandlePosts>>, TError = ErrorType<unknown>>(
 handle: string,
    params?: GetUsersHandlePostsParams, options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getUsersHandlePosts>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey = queryOptions?.queryKey ?? getGetUsersHandlePostsQueryKey(handle,params);

  

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getUsersHandlePosts>>> = ({ signal }) => getUsersHandlePosts(handle,params, requestOptions, signal);

  const query = useQuery<Awaited<ReturnType<typeof getUsersHandlePosts>>, TError, TData>(queryKey, queryFn, {enabled: !!(handle), ...queryOptions}) as  UseQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryKey;

  return query;
}


export const getPostsPostIDReplies = (
    postID: number,
    params?: GetPostsPostIDRepliesParams,
 options?: SecondParameter<typeof customInstance>,signal?: AbortSignal
) => {
      return customInstance<Posts>(
      {url: `/posts/${postID}/replies`, method: 'get',
        params, signal
    },
      options);
    }
  

export const getGetPostsPostIDRepliesQueryKey = (postID: number,
    params?: GetPostsPostIDRepliesParams,) => [`/posts/${postID}/replies`, ...(params ? [params]: [])];

    
export type GetPostsPostIDRepliesQueryResult = NonNullable<Awaited<ReturnType<typeof getPostsPostIDReplies>>>
export type GetPostsPostIDRepliesQueryError = ErrorType<unknown>

export const useGetPostsPostIDReplies = <TData = Awaited<ReturnType<typeof getPostsPostIDReplies>>, TError = ErrorType<unknown>>(
 postID: number,
    params?: GetPostsPostIDRepliesParams, options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getPostsPostIDReplies>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey = queryOptions?.queryKey ?? getGetPostsPostIDRepliesQueryKey(postID,params);

  

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getPostsPostIDReplies>>> = ({ signal }) => getPostsPostIDReplies(postID,params, requestOptions, signal);

  const query = useQuery<Awaited<ReturnType<typeof getPostsPostIDReplies>>, TError, TData>(queryKey, queryFn, {enabled: !!(postID), ...queryOptions}) as  UseQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryKey;

  return query;
}


export const postPostsPostIDReplies = (
    postID: number,
    postBodyBody: PostBodyBody,
 options?: SecondParameter<typeof customInstance>,) => {
      return customInstance<Post>(
      {url: `/posts/${postID}/replies`, method: 'post',
      headers: {'Content-Type': 'application/json', },
      data: postBodyBody
    },
      options);
    }
  


    export type PostPostsPostIDRepliesMutationResult = NonNullable<Awaited<ReturnType<typeof postPostsPostIDReplies>>>
    export type PostPostsPostIDRepliesMutationBody = PostBodyBody
    export type PostPostsPostIDRepliesMutationError = ErrorType<BadRequestResponse | UnauthorizedResponse | ForbiddenResponse | NotFoundResponse>

    export const usePostPostsPostIDReplies = <TError = ErrorType<BadRequestResponse | UnauthorizedResponse | ForbiddenResponse | NotFoundResponse>,
    
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof postPostsPostIDReplies>>, TError,{postID: number;data: PostBodyBody}, TContext>, request?: SecondParameter<typeof customInstance>}
) => {
      const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof postPostsPostIDReplies>>, {postID: number;data: PostBodyBody}> = (props) => {
          const {postID,data} = props ?? {};

          return  postPostsPostIDReplies(postID,data,requestOptions)
        }

      return useMutation<Awaited<ReturnType<typeof postPostsPostIDReplies>>, TError, {postID: number;data: PostBodyBody}, TContext>(mutationFn, mutationOptions)
    }
    
export const putPostsPostIDLikes = (
    postID: number,
 options?: SecondParameter<typeof customInstance>,) => {
      return customInstance<Likes>(
      {url: `/posts/${postID}/likes`, method: 'put'
    },
      options);
    }
  


    export type PutPostsPostIDLikesMutationResult = NonNullable<Awaited<ReturnType<typeof putPostsPostIDLikes>>>
    
    export type PutPostsPostIDLikesMutationError = ErrorType<UnauthorizedResponse | ForbiddenResponse | NotFoundResponse | Error>

    export const usePutPostsPostIDLikes = <TError = ErrorType<UnauthorizedResponse | ForbiddenResponse | NotFoundResponse | Error>,
    
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof putPostsPostIDLikes>>, TError,{postID: number}, TContext>, request?: SecondParameter<typeof customInstance>}
) => {
      const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof putPostsPostIDLikes>>, {postID: number}> = (props) => {
          const {postID} = props ?? {};

          return  putPostsPostIDLikes(postID,requestOptions)
        }

      return useMutation<Awaited<ReturnType<typeof putPostsPostIDLikes>>, TError, {postID: number}, TContext>(mutationFn, mutationOptions)
    }
    
export const deletePostsPostIDLikes = (
    postID: number,
 options?: SecondParameter<typeof customInstance>,) => {
      return customInstance<Likes>(
      {url: `/posts/${postID}/likes`, method: 'delete'
    },
      options);
    }
  


    export type DeletePostsPostIDLikesMutationResult = NonNullable<Awaited<ReturnType<typeof deletePostsPostIDLikes>>>
    
    export type DeletePostsPostIDLikesMutationError = ErrorType<UnauthorizedResponse | ForbiddenResponse | NotFoundResponse>

    export const useDeletePostsPostIDLikes = <TError = ErrorType<UnauthorizedResponse | ForbiddenResponse | NotFoundResponse>,
    
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof deletePostsPostIDLikes>>, TError,{postID: number}, TContext>, request?: SecondParameter<typeof customInstance>}
) => {
      const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof deletePostsPostIDLikes>>, {postID: number}> = (props) => {
          const {postID} = props ?? {};

          return  deletePostsPostIDLikes(postID,requestOptions)
        }

      return useMutation<Awaited<ReturnType<typeof deletePostsPostIDLikes>>, TError, {postID: number}, TContext>(mutationFn, mutationOptions)
    }
    
