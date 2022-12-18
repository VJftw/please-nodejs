import localforage from "localforage";
// import { matchSorter } from "match-sorter";
// import sortBy from "sort-by";

export type Contact = {
    id: string;
    createdAt: number;
    favourite: boolean;
    name: string;
    avatar: string;
    notes: string;
}

export function isContact(contact: unknown): contact is Contact {
    return (contact as Contact) !== undefined;
  }

export function isContacts(contacts: unknown): contacts is Contact[] {
    return (contacts as Contact[]) !== undefined;
  }


export async function getContacts(): Promise<Contact[]> {
  await fakeNetwork(`getContacts`);
  let contacts = await localforage.getItem<Contact[]>("contacts");
  if (!contacts) contacts = [];

  return contacts;

//   if (query) {
//     contacts = matchSorter(contacts, query, { keys: ["first", "last"] });
//   }
//   return contacts.sort(sortBy("last", "createdAt"));
}

export async function createContact() {
  await fakeNetwork("");
  let id = Math.random().toString(36).substring(2, 9);
  let contact = {
    id: id,
    createdAt: Date.now(),
    favourite: false,
    name: "",
    avatar: "",
    notes: "",
  };
  let contacts = await getContacts();
  contacts.unshift(contact);
  await set(contacts);
  return contact;
}

export async function getContact(id: string): Promise<Contact|null> {
  await fakeNetwork(`contact:${id}`);
  let contacts = await localforage.getItem<Contact[]>("contacts");
  if (contacts) {
    let contact = contacts.find(contact => contact.id === id);
    if (contact != undefined) {
        return contact
    }
  }

  return null;
}

export async function updateContact(id: string, updates: Contact): Promise<Contact|undefined> {
  await fakeNetwork("");
  let contacts = await localforage.getItem<Contact[]>("contacts");
  if (contacts) {
    let contact = contacts.find(contact => contact.id === id);
    if (!contact) throw new Error(`No contact found for ${id}`);
    Object.assign(contact, updates);
    await set(contacts);

    return contact;
  }

  return undefined
}

export async function deleteContact(id: string): Promise<boolean> {
  let contacts = await localforage.getItem<Contact[]>("contacts");
  if (contacts) {
    let index = contacts.findIndex(contact => contact.id === id);
    if (index > -1) {
        contacts.splice(index, 1);
        await set(contacts);
        return true;
    }
  }
  return false;
}

function set(contacts: Contact[]) {
  return localforage.setItem("contacts", contacts);
}

// fake a cache so we don't slow down stuff we've already seen
let fakeCache = new Map<string, boolean>();

async function fakeNetwork(key: string) {
  if (fakeCache.get(key)) {
    return;
  }

  fakeCache.set(key, true)
  // return new Promise(res => {
  //   setTimeout(res, Math.random() * 800);
  // });
}
