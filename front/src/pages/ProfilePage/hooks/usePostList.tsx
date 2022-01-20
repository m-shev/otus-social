import {useCallback, useEffect, useRef, useState} from 'react';
import {HttpStatus, IUseRequest, Post} from '../../../types';
import {useRequest} from '../../../hooks';
import {getPostList} from '../../../api';
import {useUserId} from './useUserId';
const TAKE = 10;
export interface IUsePostList extends Pick<IUseRequest, 'isFetch' | 'error'> {
    list: Post[];
    loadNext: () => Promise<void>;
}

const useInitialLoad = (load: () => Promise<void>): void => {
    useEffect(() => {
        // noinspection JSIgnoredPromiseFromCall
        load();
        // eslint-disable-next-line
    }, []);
};

export const usePostList = (): IUsePostList => {
    const {isFetch, error, setIsFetch, setError} = useRequest();
    const [list, setList] = useState<Post[]>([]);
    const id = useUserId();
    const {current: query} = useRef({take: TAKE, skip: 0, authorId: id});

    const loadNext = useCallback(async () => {
        setIsFetch(true);

        try {
            const resp = await getPostList(query);

            if (resp.status === HttpStatus.Ok) {
                const posts = await resp.json();
                setList([...list, ...posts]);
                query.skip += posts.length;
            } else {
                throw new Error(await resp.text());
            }
        } catch (e) {
            setError(e as Error);
        } finally {
            setIsFetch(false);
        }
    }, [list, query, setError, setIsFetch]);

    useInitialLoad(loadNext);

    return {
        isFetch,
        list,
        error,
        loadNext,
    };
};
