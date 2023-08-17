create or replace procedure clear_seed()
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
		raise notice 'admin user not found';
	else
		delete from public.users u where u.id = admin.id;
		raise notice 'admin user deleted';
	end if;
end;$$

call clear_seed()
