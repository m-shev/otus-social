import React from 'react';
import styles from './Field.module.scss';
import {FieldProps, FieldType} from './types';
import {SpecificField} from './SpecificField';

const renderCommonLabel = (props: FieldProps) => {
    return props.type !== FieldType.Checkbox;
};

export const Field: React.FC<FieldProps> = (props) => {
    const {error, ...restProps} = props;

    return (
        <div className={styles.root}>
            {renderCommonLabel(restProps) && (
                <label htmlFor={restProps.id} className={styles.label}>
                    {props.title}
                </label>
            )}

            <SpecificField {...restProps} />
            {error && <span className={styles.error}>{error}</span>}
        </div>
    );
};
