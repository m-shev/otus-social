export enum UserGender {
    Male = 'male',
    Female = 'female',
}

export type Interest = {
    id: string;
    name: string;
};

export type User = {
    id: number;
    name: string;
    surname: string;
    avatar: string;
    age: number;
    gender: UserGender;
    city: string;
    email: string;
    interests: Interest[];
};

export type UserProfile = Omit<User, 'email'> & {
    friends: UserFriend[];
};

export type CreateUserForm = Omit<User, 'id' | 'interests'> & {
    password: string;
    interests: string[];
};

export type LoginForm = {
    login: string;
    password: string;
};

export type FriendForm = {
    userId: number;
    friendId: number;
};

export type FindUserForm = {
    name?: string;
    surname?: string;
    skip: number;
    take: number;
};

export type UserFriend = Pick<User, 'id' | 'avatar' | 'name' | 'surname'>;
