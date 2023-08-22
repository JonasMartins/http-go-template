create or replace procedure seed()
language plpgsql
as $$
declare
	admin public.users%rowtype;
begin
	select *
	from public.users u
	into admin
	where u.email = 'admin@email.com';

	if not found then
		insert into public.users (uuid, name, email, "password")
		values ('4a9b3fd5-6813-4c75-9598-5fd9ae202d88','Admin', 'admin@email.com', '$2a$10$zYC48a1doguo1VoCqbmQBezAUQJKVSbGnHgoPWInNFn2idbPABUoe');
		raise notice 'admin user was sown';
	else
		raise notice 'user: % already exist', admin.name;
	end if;
end;$$

