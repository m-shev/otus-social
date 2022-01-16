export type Post = {
    id: number;
    authorId: string;
    content: string;
    imageLink?: string;
    createAt: Date;
    updateAt: Date;
};

export type CreatePostForm = Pick<Post, 'authorId' | 'content' | 'imageLink'>;
