import { timestamp } from 'drizzle-orm/pg-core';
import { integer, pgEnum, pgTable, serial, varchar } from 'drizzle-orm/pg-core';

export const statusEnum = pgEnum('status', [
  'stalling',
  'retrieving',
  'accomplished',
  'unsuccessful',
]);
// status: statusEnum('status'),

export const illustrators = pgTable('illustrators', {
  id: serial('id').primaryKey(),
  name: varchar('name', { length: 256 }).unique(),
  pixivId: varchar('pixiv_id', { length: 256 }).unique(),
  twitterId: varchar('twitter_id', { length: 256 }).unique(),
  createdAt: timestamp('created_at').defaultNow(),
  updatedAt: timestamp('updated_at').defaultNow(),
});

export const illustrations = pgTable('illustrations', {
  id: serial('id').primaryKey(),
  title: varchar('title', { length: 256 }),
  source: varchar('source', { length: 256 }),
  file: varchar('file', { length: 256 }),
  illustratorId: integer('illustrator_id').references(() => illustrators.id),
  createdAt: timestamp('created_at').defaultNow(),
  updatedAt: timestamp('updated_at').defaultNow(),
});
