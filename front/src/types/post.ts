export type Post = {
    id: number;
    authorId: number;
    content: string;
    imageLink?: string;
    createAt: Date;
    updateAt: Date;
};

export type CreatePostForm = Pick<Post, 'authorId' | 'content' | 'imageLink'>;
