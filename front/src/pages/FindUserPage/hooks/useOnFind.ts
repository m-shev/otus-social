import {HttpStatus, IUseRequestState, UserFriend} from '../../../types';
import {useRequest} from '../../../hooks';
import {useCallback, useState} from 'react';
import {FindForm} from '../const';
import {getUserList} from '../../../api';

const DEFAULT_TAKE = 500;

export interface IUseOnFind extends IUseRequestState {
    findList: UserFriend[];
    onSubmit: (values: FindForm) => Promise<void>;
    firstSearchFlag: boolean;
}

export const useOnFind = (): IUseOnFind => {
    const {setIsFetch, setRequestState, setError, error, requestState, isFetch} = useRequest();
    const [findList, setFindList] = useState([]);
    const [firstSearchFlag, setFlag] = useState(false);

    const onSubmit = useCallback(
        async (values: FindForm) => {
            setIsFetch(true);

            try {
                const resp = await getUserList({
                    skip: 0,
                    take: DEFAULT_TAKE,
                    name: values.name,
                    surname: values.surname,
                });

                if (resp.status === HttpStatus.Ok) {
                    setFlag(true);
                    setRequestState('success');
                    setFindList(await resp.json());
                } else {
                    setRequestState('fail');
                    throw new Error(await resp.text());
                }
            } catch (e) {
                setError(e as Error);
            } finally {
                setIsFetch(false);
            }
        },
        [setError, setIsFetch, setRequestState],
    );

    return {
        findList,
        onSubmit,
        error,
        requestState,
        isFetch,
        firstSearchFlag,
    };
};
