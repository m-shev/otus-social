import {useEffect, useState} from 'react';

const MAX_DOT = 3;
const DOT = '.';
const INTERVAL = 500;

export const useDots = (): string => {
    const [dots, setDots] = useState([DOT, DOT, DOT]);

    useEffect(() => {
        const intervalId = setInterval(() => {
            if (dots.length === MAX_DOT) {
                setDots([DOT]);
            } else {
                dots.push(DOT);
                setDots([...dots]);
            }
        }, INTERVAL);

        return () => {
            return clearInterval(intervalId);
        };
    }, [dots]);

    return dots.join('');
};
