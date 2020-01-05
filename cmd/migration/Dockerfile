FROM migrate-base

COPY migration /migration
COPY Makefile /Makefile

ENV DB_USER \
    DB_PASSWORD \
    DB_HOST \
    DB_SCHEMA

ENTRYPOINT ["make"]
CMD ["migration"]