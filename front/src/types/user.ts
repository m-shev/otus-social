export enum UserGender {
    Male = 'Male',
    Female = 'Female',
}

export type Interest = {
    id: string;
    name: string;
};

export type User = {
    id: number;
    name: string;
    surname: string;
    age: number;
    city: string;
    email: string;
    interests: Interest[];
};

export type UserProfile = Omit<User, 'email'>;

export type CreateUserForm = Omit<User, 'id' | 'interests'> & {
    password: string;
    interests: string[];
};

export type LoginForm = {
    login: string;
    password: string;
};
