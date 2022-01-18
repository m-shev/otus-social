import {HttpStatus, IUseRequestState} from '../../../types';
import {useRequest, useUser} from '../../../hooks';
import {useCallback} from 'react';
import {CreatePostField} from '../const';
import {createPostPost} from '../../../api';

export interface IUseCreatePost
    extends Pick<IUseRequestState, 'isFetch' | 'error' | 'requestState'> {
    onCreatePost: (form: CreatePostField) => Promise<void>;
}

export const useCreatePost = (): IUseCreatePost => {
    const {setIsFetch, setRequestState, setError, error, isFetch, requestState} = useRequest();
    const user = useUser();

    // noinspection ExceptionCaughtLocallyJS
    const onCreatePost = useCallback(
        async (values: CreatePostField) => {
            setIsFetch(true);
            console.log('dev -------->', 'start request');
            try {
                const {content, imageLink} = values;

                if (user === null) {
                    // noinspection ExceptionCaughtLocallyJS
                    throw new Error('user must be not null');
                }

                const resp = await createPostPost({
                    content,
                    imageLink,
                    authorId: user.id,
                });

                if (resp.status !== HttpStatus.Ok) {
                    // noinspection ExceptionCaughtLocallyJS
                    throw new Error(await resp.text());
                }

                setRequestState('success');
            } catch (e) {
                setError(e as Error);
                setRequestState('fail');
            } finally {
                setIsFetch(false);
            }
        },
        [setError, setIsFetch, setRequestState, user],
    );

    return {
        error,
        isFetch,
        onCreatePost,
        requestState,
    };
};
