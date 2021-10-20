import {useState} from 'react';
import {IUseRequest, RequestState} from '../types';

export const useRequest = (): IUseRequest => {
    const [isFetch, setIsFetch] = useState(false);
    const [error, setError] = useState<Error | null>(null);
    const [requestState, setRequestState] = useState<RequestState>(null);

    return {
        isFetch,
        setIsFetch,
        error,
        setError,
        requestState,
        setRequestState,
    };
};
