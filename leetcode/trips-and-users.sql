SELECT t.request_at AS Day,
  CAST(
    CAST(SUM(
      CASE
        WHEN t.status = 'cancelled_by_driver' OR t.status = 'cancelled_by_client'
        THEN 1
        ELSE 0
      END
    ) AS DECIMAL(10, 2)
  ) / CAST(COUNT(t.id) AS DECIMAL(10, 2)) AS DECIMAL(10, 2)) AS 'Cancellation Rate'
FROM Trips AS t
JOIN Users AS u1
  ON t.client_id = u1.users_id
JOIN Users AS u2
  ON T.driver_id = u2.users_id
WHERE
  u1.banned = 'No' AND u2.banned = 'No' AND t.request_at BETWEEN '2013-10-01' AND '2013-10-03'
GROUP BY
  t.request_at;
