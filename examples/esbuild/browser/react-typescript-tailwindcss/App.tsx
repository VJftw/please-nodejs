import React, { useCallback, useState } from "react";
import ReactDOM from "react-dom";
import {
  createHashRouter,
  RouterProvider,
  Link,
  isRouteErrorResponse,
  useRouteError,
  Outlet,
  useLoaderData,
  Form,
} from "react-router-dom";

import ContactComponent, { loader as contactLoader } from "./contacts/Contact";
import EditContact from "./contacts/EditContact";
import {
  editAction,
  deleteAction
} from "./contacts/actions";
import { Contact, getContacts, createContact } from "./contacts/data";

const router = createHashRouter([
  {
    path: "/",
    element: <Root />,
    errorElement: <ErrorPage />,
    loader: rootLoader,
    action: rootAction,
    children: [
      {
        path: "contacts/:contactId",
        element: <ContactComponent />,
        loader: contactLoader,
      },
      {
        path: "contacts/:contactId/edit",
        element: <EditContact />,
        loader: contactLoader,
        action: editAction,
      },
      {
        path: "contacts/:contactId/delete",
        action: deleteAction,
      },
    ],
  },
]);

function isContacts(contacts: unknown): contacts is Contact[] {
  return (contacts as Contact[]) !== undefined;
}

export default function Root() {
  const [contacts, setContacts] = useState<Contact[]>([]);
  const loadedContacts = useLoaderData();

  if (isContacts(loadedContacts)) {
    if (loadedContacts != contacts) {
      setContacts(loadedContacts);
    }
  }

  return (
    <>
      <div className="flex xs:flex-col sm:flex-row">

      <div className="max-w-xs bg-gray-50">
        <div className="flex flex-col justify-between h-screen bg-white border-r">
          <div className="px-4 py-6">
            <span className="block w-32 h-10 bg-gray-200 rounded-lg"></span>

            <nav aria-label="Main Nav" className="flex flex-col mt-6 space-y-1">

              <details
                className="group [&_summary::-webkit-details-marker]:hidden"
                open
              >
                <summary className="flex items-center px-4 py-2 text-gray-500 rounded-lg cursor-pointer hover:bg-gray-100 hover:text-gray-700">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    className="w-5 h-5 opacity-75"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                    stroke-width="2"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"
                    />
                  </svg>

                  <span className="ml-3 text-sm font-medium"> Contacts </span>

                  <span className="ml-auto transition duration-300 shrink-0 group-open:-rotate-180">
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      className="w-5 h-5"
                      viewBox="0 0 20 20"
                      fill="currentColor"
                    >
                      <path
                        fill-rule="evenodd"
                        d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z"
                        clip-rule="evenodd"
                      />
                    </svg>
                  </span>
                </summary>

                <nav
                  aria-label="Contacts Nav"
                  className="mt-1.5 ml-8 flex flex-col"
                >
                  <div className="flex items-center text-gray-700 rounded-lg">

                  <Form method="post" className="flex items-center w-full">
                    <button type="submit" className="cursor-pointer flex w-full p-1 items-center rounded border border-indigo-600 text-indigo-600 hover:bg-indigo-600 hover:text-white focus:outline-none focus:ring active:bg-indigo-500"
                      ><span className="text-sm text-center font-medium w-full">New Contact</span></button>
                  </Form>
                  </div>
                  {contacts.length ? (
                    <ul>
                      {contacts.map((contact) => (
                        <li key={contact.id}>
                          <Link
                            to={`contacts/${contact.id}`}
                            className="flex items-center px-4 py-2 text-gray-500 rounded-lg hover:bg-gray-100 hover:text-gray-700"
                          >
                            <svg
                              xmlns="http://www.w3.org/2000/svg"
                              className="w-5 h-5 opacity-75"
                              fill="none"
                              viewBox="0 0 24 24"
                              stroke="currentColor"
                              stroke-width="2"
                            >
                              <path
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                d="M10 6H5a2 2 0 00-2 2v9a2 2 0 002 2h14a2 2 0 002-2V8a2 2 0 00-2-2h-5m-4 0V5a2 2 0 114 0v1m-4 0a2 2 0 104 0m-5 8a2 2 0 100-4 2 2 0 000 4zm0 0c1.306 0 2.417.835 2.83 2M9 14a3.001 3.001 0 00-2.83 2M15 11h3m-3 4h2"
                              />
                            </svg>
                            <span className="ml-3 text-sm font-medium">
                              {contact.name != "" ? (
                                <>{contact.name}</>
                              ) : (
                                <i>No Name</i>
                              )}{" "}
                              {contact.favourite && <span>★</span>}
                            </span>
                          </Link>
                        </li>
                      ))}
                    </ul>
                  ) : (
                    <p>No contacts</p>
                  )}
                </nav>
              </details>
            </nav>
          </div>
        </div>
      </div>

      <div className="flex w-full"><Outlet /></div>

      </div>
    </>
  );
}

export function ErrorPage() {
  const error = useRouteError();
  console.error(error);

  if (isRouteErrorResponse(error)) {
    return (
      <div id="error-page">
        <h1>Oops!</h1>
        <p>Sorry, an unexpected error has occurred.</p>
        <p>
          <i>
            {error.status}: {error.statusText}
          </i>
        </p>
      </div>
    );
  }

  return <></>;
}

export async function rootAction(): Promise<Contact> {
  const contact = await createContact();

  return contact;
}

export async function rootLoader(): Promise<Contact[]> {
  const contacts = await getContacts();
  console.log(contacts);
  return contacts;
}

ReactDOM.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
  document.getElementById("root")
);
