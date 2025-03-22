-- Convert existing payment_type values to new numeric format
UPDATE transactions 
SET payment_type = CASE 
    WHEN payment_type = 'cach' THEN '1'
    WHEN payment_type = 'credit_card' THEN '2'
    ELSE '1'
END;

-- Change column type from TEXT to INT8
ALTER TABLE transactions 
ALTER COLUMN payment_type TYPE INT8 
USING payment_type::INT8;
