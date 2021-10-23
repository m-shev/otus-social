import {useMemo} from 'react';
import {Interest} from '../../../types';

export const useInterests = (interestList: Interest[] = []): string | null => {
    return useMemo(() => {
        if (interestList.length) {
            return interestList
                .map((item) => {
                    return item.name;
                })
                .join(', ');
        }

        return null;
    }, [interestList]);
};
