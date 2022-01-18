import {useHistory} from 'react-router';
import {useEffect} from 'react';
import {useUser} from './useUser';

export interface ICondition {
    (): boolean;
}

export const useRedirectToProfile = (conditionFunc?: ICondition): void => {
    const user = useUser();
    const history = useHistory();

    useEffect(() => {
        if (user != null) {
            if ((conditionFunc && conditionFunc()) || !conditionFunc) {
                history.push(`profile?id=${user.id}`);
            }
        }

        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [user, conditionFunc]);
};
