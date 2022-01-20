import {HttpStatus, IUseRequestState, LoginForm} from '../../../types';
import {useRequest} from '../../../hooks';
import {login} from '../../../api';
import {useCallback} from 'react';
import {userAuthEvent} from '../../../store/user';

export interface IUseLogin extends IUseRequestState {
    onSubmit: (values: LoginForm) => void;
}

export const useOnLogin = (): IUseLogin => {
    const {error, setError, isFetch, setIsFetch, setRequestState, requestState} = useRequest();

    const onSubmit = useCallback(
        async (values: LoginForm): Promise<void> => {
            setIsFetch(true);
            const resp = await login(values);

            if (resp.status === HttpStatus.Ok) {
                setRequestState('success');
                userAuthEvent(await resp.json());
            } else {
                setRequestState('fail');
                const errorText = await resp.text();
                setError(new Error(errorText));
            }

            setIsFetch(false);
        },
        [setError, setIsFetch, setRequestState],
    );

    return {
        error,
        isFetch,
        requestState,
        onSubmit,
    };
};
